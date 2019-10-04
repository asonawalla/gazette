// +build integration

package integration

import (
	"testing"
	"time"

	"github.com/jgraettinger/urkel"
)

func TestCrashOneEtcdMember(t *testing.T) {
	var pod = urkel.FetchPods(t, testNamespace, etcdPodSelector)[0]

	var fs = urkel.NewFaultSet(t)
	defer fs.RemoveAll()

	fs.Crash(pod)
	time.Sleep(time.Second * 30)
}

func TestCrashTwoBrokers(t *testing.T) {
	var pods = urkel.FetchPods(t, testNamespace, brokerPodSelector)[:2]

	var fs = urkel.NewFaultSet(t)
	defer fs.RemoveAll()

	fs.Crash(pods...)
	time.Sleep(time.Second * 30)
}

func TestCrashOneSummer(t *testing.T) {
	var pod = urkel.FetchPods(t, testNamespace, summerPodSelector)[0]

	var fs = urkel.NewFaultSet(t)
	defer fs.RemoveAll()

	fs.Crash(pod)
	time.Sleep(time.Second * 30)
}

func TestDeleteAllBrokers(t *testing.T) {
	var pods = urkel.FetchPods(t, testNamespace, brokerPodSelector)

	var fs = urkel.NewFaultSet(t)
	defer fs.RemoveAll()

	fs.Delete(pods...)
	time.Sleep(time.Second * 30)
}

func TestDeleteAllSummers(t *testing.T) {
	var pods = urkel.FetchPods(t, testNamespace, summerPodSelector)

	var fs = urkel.NewFaultSet(t)
	defer fs.RemoveAll()

	fs.Delete(pods...)
	time.Sleep(time.Second * 30)
}
