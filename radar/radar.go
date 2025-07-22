package radar

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"sync"
	"unsafe"

	pipe "github.com/b4b4r07/go-pipe"
	"github.com/gin-gonic/gin"
	onvif "github.com/ugparu/onvif/lib"

	"github.com/sirupsen/logrus"
)

func GetNetInterfaces() ([]string, error) {
	var b bytes.Buffer
	pipe.Command(&b,
		exec.Command("ls", "/sys/class/net/"),
		exec.Command("grep", "-v", `lo\|docker`),
		exec.Command("sed", `$!s/$/,/`),
	)

	interfaces := strings.Split(b.String(), ",")
	logrus.Debug(fmt.Sprintf("Net Interfaces: %v", interfaces))

	return interfaces, nil
}

func removeEmpties(s []string) []string {
	r := []string{}
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func GetOnvifCameras() ([]string, error) {
	netInts, err := GetNetInterfaces()
	if err != nil {
		return []string{}, err
	}

	type CamsIPs [1024]string
	type Nics map[string]CamsIPs

	nics := make(Nics, 1024)

	for i := 0; i < len(netInts); i++ {
		nics[netInts[i]] = [1024]string{}
	}

	var wg sync.WaitGroup
	for k := range nics {
		wg.Add(1)
		go func(nic *Nics, nicName string) {
			defer wg.Done()
			discovery, err := onvif.GetAvailableDevicesAtSpecificEthernetInterface(strings.TrimSpace(nicName))
			if err != nil {
				logrus.Debug(fmt.Sprintf("Continued interface: %s", nicName))
				logrus.Debug(fmt.Sprintf("Err: %s", err))
				return
			}

			log.Printf("disc on nic %s %#v", nicName, discovery)
			if len(discovery) > 0 {
				ipBuf := [1024]string{}
				for j := 0; j < len(discovery); j++ {

					rs := reflect.ValueOf(&discovery[j]).Elem()
					rxaddr := rs.Field(0).Field(0)
					rxaddr = reflect.NewAt(rxaddr.Type(), unsafe.Pointer(rxaddr.UnsafeAddr())).Elem()
					xaddr := rxaddr.Interface().(string)
					logrus.Debug(fmt.Sprintf("Camera ip: %#v", xaddr))

					log.Printf("%s", xaddr)
					if len([]rune(xaddr)) < 1 {
						continue
					}

					if xaddr != "" {
						ipBuf[j] = xaddr
					}

				}
				(*nic)[nicName] = ipBuf
			}
			logrus.Debug(fmt.Sprintf("%s %#v", strings.TrimSpace(nicName), discovery))
		}(&nics, k)
	}
	wg.Wait()

	res := []string{}
	for _, ips := range nics {
		res = append(res, ips[:]...)
	}

	res = removeEmpties(res)
	logrus.Info(fmt.Sprintf("\n Cameras from current subnet: %#v", res))

	fakeIps := strings.Split(os.Getenv("FAKE_CAMERAS_IPS"), ",")
	if len(fakeIps) > 1 {
		logrus.Info(fmt.Sprintf("\n FAKE MODE ENABLED. Fake Cameras list: %#v", fakeIps))
		return fakeIps, nil
	}
	return res, nil
}

func GetOnvifCamerasGin(c *gin.Context) {
	res, err := GetOnvifCameras()
	if err != nil {
		c.JSON(400, gin.H{"error": fmt.Sprintf("%s", err)})
	}
	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
