package helper

import (
	. "github.com/onsi/gomega"
	"github.com/rancher/rancher/tests/framework/clients/rancher"
	"github.com/rancher/rancher/tests/framework/extensions/cloudcredentials"
	"github.com/rancher/rancher/tests/framework/extensions/cloudcredentials/google"
	"github.com/rancher/rancher/tests/framework/pkg/session"
)

type Context struct {
	CloudCred     *cloudcredentials.CloudCredential
	RancherClient *rancher.Client
	Session       *session.Session
}

func CommonBeforeEach() Context {
	testSession := session.NewSession()
	rancherClient, err := rancher.NewClient("", testSession)
	Expect(err).To(BeNil())
	cloudCredential, err := google.CreateGoogleCloudCredentials(rancherClient)
	Expect(err).To(BeNil())
	return Context{
		CloudCred:     cloudCredential,
		RancherClient: rancherClient,
		Session:       testSession,
	}
}

func CommonAfterEach(ctx Context) {

}
