/*
	Copyright NetFoundry, Inc.

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	https://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package tunnel

import (
	"github.com/openziti/foundation/util/info"
	"github.com/openziti/sdk-golang/ziti"
	"github.com/sirupsen/logrus"
	"io"
	"net"
)

var log = logrus.StandardLogger()

func Run(context ziti.Context, service string, clientConn net.Conn) {
	zitiConn, err := context.Dial(service)
	if err != nil {
		log.Errorf("zt.Dial(%s) failed: %s", service, err.Error())
		clientConn.Close()
		return
	}

	loggerFields := logrus.Fields{
		"src-remote": clientConn.RemoteAddr(), "src-local": clientConn.LocalAddr(),
		"dst-local": zitiConn.LocalAddr(), "dst-remote": zitiConn.RemoteAddr()}

	log := log.WithFields(loggerFields)
	log.Info("tunnel started")

	doneSend := make(chan int64)
	doneRecv := make(chan int64)

	go myCopy(clientConn, zitiConn, doneSend)

	go myCopy(zitiConn, clientConn, doneRecv)

	var n1, n2 int64
	for count := 2; count > 0; {
		select {
		case n1 = <-doneSend:
		case n2 = <-doneRecv:
		}
		count = count - 1
	}

	log.Infof("tunnel closed: %d bytes sent; %d bytes received", n2, n1)
}

func myCopy(dst net.Conn, src net.Conn, done chan int64) {
	loggerFields := logrus.Fields{
		"src-remote": src.RemoteAddr(), "src-local": src.LocalAddr(),
		"dst-local": dst.LocalAddr(), "dst-remote": dst.RemoteAddr()}

	defer dst.Close()
	defer log.WithFields(loggerFields).Info("stopping pipe")
	copyBuf := make([]byte, info.MaxUdpPacketSize)
	n, err := io.CopyBuffer(dst, src, copyBuf)
	done <- n

	if err != nil {
		log.WithFields(loggerFields).Error(err.Error())
	}
}
