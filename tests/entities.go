// +build apitests

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

package tests

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"github.com/openziti/edge/eid"
	"math/big"
	"sort"
	"time"

	"github.com/Jeffail/gabs"
)

type entity interface {
	getId() string
	setId(string)
	getEntityType() string
	toJson(create bool, ctx *TestContext, fields ...string) string
	validate(ctx *TestContext, c *gabs.Container)
}

type service struct {
	Id                 string
	Name               string
	terminatorStrategy string
	roleAttributes     []string
	configs            []string
	permissions        []string
	tags               map[string]interface{}
}

func (entity *service) getId() string {
	return entity.Id
}

func (entity *service) setId(id string) {
	entity.Id = id
}

func (entity *service) getEntityType() string {
	return "services"
}

func (entity *service) toJson(_ bool, ctx *TestContext, _ ...string) string {
	entityData := gabs.New()
	ctx.setJsonValue(entityData, entity.Name, "name")
	ctx.setJsonValue(entityData, entity.terminatorStrategy, "terminatorStrategy")
	ctx.setJsonValue(entityData, entity.roleAttributes, "roleAttributes")
	ctx.setJsonValue(entityData, entity.configs, "configs")

	if len(entity.tags) > 0 {
		ctx.setJsonValue(entityData, entity.tags, "tags")
	}

	return entityData.String()
}

func (entity *service) validate(ctx *TestContext, c *gabs.Container) {
	if entity.tags == nil {
		entity.tags = map[string]interface{}{}
	}
	ctx.pathEquals(c, entity.Name, path("name"))
	ctx.pathEquals(c, entity.terminatorStrategy, path("terminatorStrategy"))
	ctx.pathEquals(c, entity.tags, path("tags"))

	sort.Strings(entity.roleAttributes)
	ctx.pathEqualsStringSlice(c, entity.roleAttributes, path("roleAttributes"))

	sort.Strings(entity.permissions)
	ctx.pathEqualsStringSlice(c, entity.permissions, path("permissions"))
}

type terminator struct {
	id        string
	serviceId string
	routerId  string
	binding   string
	address   string
	tags      map[string]interface{}
}

func (entity *terminator) getId() string {
	return entity.id
}

func (entity *terminator) setId(id string) {
	entity.id = id
}

func (entity *terminator) getEntityType() string {
	return "terminators"
}

func (entity *terminator) toJson(_ bool, ctx *TestContext, _ ...string) string {
	entityData := gabs.New()
	ctx.setJsonValue(entityData, entity.serviceId, "service")
	ctx.setJsonValue(entityData, entity.routerId, "router")
	ctx.setJsonValue(entityData, entity.binding, "binding")
	ctx.setJsonValue(entityData, entity.address, "address")

	if len(entity.tags) > 0 {
		ctx.setJsonValue(entityData, entity.tags, "tags")
	}

	return entityData.String()
}

func (entity *terminator) validate(ctx *TestContext, c *gabs.Container) {
	if entity.tags == nil {
		entity.tags = map[string]interface{}{}
	}
	ctx.pathEquals(c, entity.serviceId, path("serviceId"))
	ctx.pathEquals(c, entity.routerId, path("routerId"))
	ctx.pathEquals(c, entity.binding, path("binding"))
	ctx.pathEquals(c, entity.address, path("address"))
	ctx.pathEquals(c, entity.tags, path("tags"))
}

func newTestIdentity(isAdmin bool, roleAttributes ...string) *identity {
	return &identity{
		name:           eid.New(),
		identityType:   "User",
		isAdmin:        isAdmin,
		roleAttributes: roleAttributes,
	}
}

type identity struct {
	Id             string
	name           string
	identityType   string
	isAdmin        bool
	enrollment     map[string]interface{}
	roleAttributes []string
	tags           map[string]interface{}
}

func (entity *identity) getId() string {
	return entity.Id
}

func (entity *identity) setId(id string) {
	entity.Id = id
}

func (entity *identity) getEntityType() string {
	return "identities"
}

func (entity *identity) toJson(isCreate bool, ctx *TestContext, _ ...string) string {
	entityData := gabs.New()
	ctx.setJsonValue(entityData, entity.name, "name")
	ctx.setJsonValue(entityData, entity.identityType, "type")
	ctx.setJsonValue(entityData, entity.isAdmin, "isAdmin")
	ctx.setJsonValue(entityData, entity.enrollment, "enrollment")
	ctx.setJsonValue(entityData, entity.roleAttributes, "roleAttributes")

	if isCreate {
		if entity.enrollment == nil {
			enrollments := map[string]interface{}{
				"updb": entity.name,
			}
			ctx.setJsonValue(entityData, enrollments, "enrollment")
		}
	}

	ctx.setJsonValue(entityData, entity.tags, "tags")

	return entityData.String()
}
func (entity *identity) validate(ctx *TestContext, c *gabs.Container) {
	if entity.tags == nil {
		entity.tags = map[string]interface{}{}
	}
	ctx.pathEquals(c, entity.name, path("name"))
	sort.Strings(entity.roleAttributes)
	ctx.pathEqualsStringSlice(c, entity.roleAttributes, path("roleAttributes"))
	ctx.pathEquals(c, entity.tags, path("tags"))
}

func newTestEdgeRouter(roleAttributes ...string) *edgeRouter {
	return &edgeRouter{
		name:           eid.New(),
		roleAttributes: roleAttributes,
	}
}

type edgeRouter struct {
	id             string
	name           string
	roleAttributes []string
	tags           map[string]interface{}
}

func (entity *edgeRouter) getId() string {
	return entity.id
}

func (entity *edgeRouter) setId(id string) {
	entity.id = id
}

func (entity *edgeRouter) getEntityType() string {
	return "edge-routers"
}

func (entity *edgeRouter) toJson(_ bool, ctx *TestContext, _ ...string) string {
	entityData := gabs.New()
	ctx.setJsonValue(entityData, entity.name, "name")
	ctx.setJsonValue(entityData, entity.roleAttributes, "roleAttributes")

	ctx.setJsonValue(entityData, entity.tags, "tags")

	return entityData.String()
}

func (entity *edgeRouter) validate(ctx *TestContext, c *gabs.Container) {
	if entity.tags == nil {
		entity.tags = map[string]interface{}{}
	}
	ctx.pathEquals(c, entity.name, path("name"))
	sort.Strings(entity.roleAttributes)
	ctx.pathEqualsStringSlice(c, entity.roleAttributes, path("roleAttributes"))
	ctx.pathEquals(c, entity.tags, path("tags"))
}

func newEdgeRouterPolicy(semantic *string, edgeRouterRoles, identityRoles []string) *edgeRouterPolicy {
	return &edgeRouterPolicy{
		name:            eid.New(),
		semantic:        semantic,
		edgeRouterRoles: edgeRouterRoles,
		identityRoles:   identityRoles,
	}
}

type edgeRouterPolicy struct {
	id              string
	name            string
	semantic        *string
	edgeRouterRoles []string
	identityRoles   []string
	tags            map[string]interface{}
}

func (entity *edgeRouterPolicy) getId() string {
	return entity.id
}

func (entity *edgeRouterPolicy) setId(id string) {
	entity.id = id
}

func (entity *edgeRouterPolicy) getEntityType() string {
	return "edge-router-policies"
}

func (entity *edgeRouterPolicy) toJson(_ bool, ctx *TestContext, _ ...string) string {
	entityData := gabs.New()
	ctx.setJsonValue(entityData, entity.name, "name")
	if entity.semantic != nil {
		ctx.setJsonValue(entityData, *entity.semantic, "semantic")
	}
	ctx.setJsonValue(entityData, entity.edgeRouterRoles, "edgeRouterRoles")
	ctx.setJsonValue(entityData, entity.identityRoles, "identityRoles")

	if len(entity.tags) > 0 {
		ctx.setJsonValue(entityData, entity.tags, "tags")
	}
	return entityData.String()
}

func (entity *edgeRouterPolicy) validate(ctx *TestContext, c *gabs.Container) {
	if entity.tags == nil {
		entity.tags = map[string]interface{}{}
	}
	if entity.semantic == nil {
		t := "AllOf"
		entity.semantic = &t
	}
	ctx.pathEquals(c, entity.name, path("name"))
	ctx.pathEquals(c, *entity.semantic, path("semantic"))
	sort.Strings(entity.edgeRouterRoles)
	ctx.pathEqualsStringSlice(c, entity.edgeRouterRoles, path("edgeRouterRoles"))
	sort.Strings(entity.identityRoles)
	ctx.pathEqualsStringSlice(c, entity.identityRoles, path("identityRoles"))
	ctx.pathEquals(c, entity.tags, path("tags"))
}

func newServiceEdgeRouterPolicy(semantic *string, edgeRouterRoles, serviceRoles []string) *serviceEdgeRouterPolicy {
	return &serviceEdgeRouterPolicy{
		name:            eid.New(),
		semantic:        semantic,
		edgeRouterRoles: edgeRouterRoles,
		serviceRoles:    serviceRoles,
	}
}

type serviceEdgeRouterPolicy struct {
	id              string
	name            string
	semantic        *string
	edgeRouterRoles []string
	serviceRoles    []string
	tags            map[string]interface{}
}

func (entity *serviceEdgeRouterPolicy) getId() string {
	return entity.id
}

func (entity *serviceEdgeRouterPolicy) setId(id string) {
	entity.id = id
}

func (entity *serviceEdgeRouterPolicy) getEntityType() string {
	return "service-edge-router-policies"
}

func (entity *serviceEdgeRouterPolicy) toJson(_ bool, ctx *TestContext, _ ...string) string {
	entityData := gabs.New()
	ctx.setJsonValue(entityData, entity.name, "name")
	if entity.semantic != nil {
		ctx.setJsonValue(entityData, *entity.semantic, "semantic")
	}
	ctx.setJsonValue(entityData, entity.edgeRouterRoles, "edgeRouterRoles")
	ctx.setJsonValue(entityData, entity.serviceRoles, "serviceRoles")

	if len(entity.tags) > 0 {
		ctx.setJsonValue(entityData, entity.tags, "tags")
	}
	return entityData.String()
}

func (entity *serviceEdgeRouterPolicy) validate(ctx *TestContext, c *gabs.Container) {
	if entity.tags == nil {
		entity.tags = map[string]interface{}{}
	}
	if entity.semantic == nil {
		t := "AllOf"
		entity.semantic = &t
	}
	ctx.pathEquals(c, entity.name, path("name"))
	ctx.pathEquals(c, *entity.semantic, path("semantic"))
	sort.Strings(entity.edgeRouterRoles)
	ctx.pathEqualsStringSlice(c, entity.edgeRouterRoles, path("edgeRouterRoles"))
	sort.Strings(entity.serviceRoles)
	ctx.pathEqualsStringSlice(c, entity.serviceRoles, path("serviceRoles"))
	ctx.pathEquals(c, entity.tags, path("tags"))
}

func newServicePolicy(policyType string, semantic *string, serviceRoles, identityRoles []string) *servicePolicy {
	return &servicePolicy{
		name:          eid.New(),
		policyType:    policyType,
		semantic:      semantic,
		serviceRoles:  serviceRoles,
		identityRoles: identityRoles,
	}
}

type servicePolicy struct {
	id            string
	name          string
	policyType    string
	semantic      *string
	identityRoles []string
	serviceRoles  []string
	tags          map[string]interface{}
}

func (entity *servicePolicy) getId() string {
	return entity.id
}

func (entity *servicePolicy) setId(id string) {
	entity.id = id
}

func (entity *servicePolicy) getEntityType() string {
	return "service-policies"
}

func (entity *servicePolicy) toJson(_ bool, ctx *TestContext, _ ...string) string {
	entityData := gabs.New()
	ctx.setJsonValue(entityData, entity.name, "name")
	ctx.setJsonValue(entityData, entity.policyType, "type")
	if entity.semantic != nil {
		ctx.setJsonValue(entityData, entity.semantic, "semantic")
	}
	ctx.setJsonValue(entityData, entity.identityRoles, "identityRoles")
	ctx.setJsonValue(entityData, entity.serviceRoles, "serviceRoles")

	if len(entity.tags) > 0 {
		ctx.setJsonValue(entityData, entity.tags, "tags")
	}
	return entityData.String()
}

func (entity *servicePolicy) validate(ctx *TestContext, c *gabs.Container) {
	if entity.tags == nil {
		entity.tags = map[string]interface{}{}
	}
	if entity.semantic == nil {
		t := "AllOf"
		entity.semantic = &t
	}
	ctx.pathEquals(c, entity.name, path("name"))
	ctx.pathEquals(c, entity.policyType, path("type"))
	ctx.pathEquals(c, *entity.semantic, path("semantic"))
	sort.Strings(entity.identityRoles)
	ctx.pathEqualsStringSlice(c, entity.identityRoles, path("identityRoles"))
	sort.Strings(entity.serviceRoles)
	ctx.pathEqualsStringSlice(c, entity.serviceRoles, path("serviceRoles"))
	ctx.pathEquals(c, entity.tags, path("tags"))
}

type Config struct {
	Id           string
	ConfigTypeId string
	Name         string
	Data         map[string]interface{}
	Tags         map[string]interface{}
	sendType     bool
}

func (entity *Config) getId() string {
	return entity.Id
}

func (entity *Config) setId(id string) {
	entity.Id = id
}

func (entity *Config) getEntityType() string {
	return "configs"
}

func (entity *Config) toJson(isCreate bool, ctx *TestContext, fields ...string) string {
	entityData := gabs.New()
	ctx.setValue(entityData, entity.Name, fields, "name")
	if isCreate || entity.sendType {
		ctx.setValue(entityData, entity.ConfigTypeId, fields, "configTypeId")
	}
	ctx.setValue(entityData, entity.Data, fields, "data")
	ctx.setValue(entityData, entity.Tags, fields, "tags")
	return entityData.String()
}

func (entity *Config) validate(ctx *TestContext, c *gabs.Container) {
	if entity.Tags == nil {
		entity.Tags = map[string]interface{}{}
	}
	ctx.pathEquals(c, entity.Name, path("name"))
	ctx.pathEquals(c, entity.ConfigTypeId, path("configTypeId"))
	ctx.pathEquals(c, entity.Data, path("data"))
	ctx.pathEquals(c, entity.Tags, path("tags"))
}

type configType struct {
	Id     string
	Name   string
	Schema map[string]interface{}
	Tags   map[string]interface{}
}

func (entity *configType) getId() string {
	return entity.Id
}

func (entity *configType) setId(id string) {
	entity.Id = id
}

func (entity *configType) getEntityType() string {
	return "config-types"
}

func (entity *configType) toJson(isCreate bool, ctx *TestContext, fields ...string) string {
	entityData := gabs.New()
	ctx.setValue(entityData, entity.Name, fields, "name")
	ctx.setValue(entityData, entity.Schema, fields, "schema")
	ctx.setValue(entityData, entity.Tags, fields, "tags")
	return entityData.String()
}

func (entity *configType) validate(ctx *TestContext, c *gabs.Container) {
	if entity.Tags == nil {
		entity.Tags = map[string]interface{}{}
	}
	ctx.pathEquals(c, entity.Name, path("name"))
	ctx.pathEquals(c, entity.Schema, path("schema"))
	ctx.pathEquals(c, entity.Tags, path("tags"))
}

type apiSession struct {
	id          string
	token       string
	identityId  string
	configTypes []string
	tags        map[string]interface{}
}

func (entity *apiSession) getId() string {
	return entity.id
}

func (entity *apiSession) setId(id string) {
	entity.id = id
}

func (entity *apiSession) getEntityType() string {
	return "apiSessions"
}

func (entity *apiSession) toJson(_ bool, ctx *TestContext, fields ...string) string {
	ctx.Req.FailNow("should not be called")
	return ""
}

func (entity *apiSession) validate(ctx *TestContext, c *gabs.Container) {
	if entity.tags == nil {
		entity.tags = map[string]interface{}{}
	}
	ctx.pathEquals(c, entity.token, path("token"))
	ctx.pathEquals(c, entity.identityId, path("identity", "id"))
	ctx.pathEquals(c, entity.configTypes, path("configTypes"))
	ctx.pathEquals(c, entity.tags, path("tags"))
}

type configValidatingService struct {
	*service
	configs map[string]*Config
}

func (entity *configValidatingService) validate(ctx *TestContext, c *gabs.Container) {
	configs := c.Path("config")
	if len(entity.configs) == 0 && configs == nil {
		return
	}

	children, err := configs.Children()
	ctx.Req.NoError(err)
	ctx.Req.Equal(len(entity.configs), len(children))
	for configType, config := range entity.configs {
		ctx.pathEquals(configs, config.Data, path(configType))
	}
}

func newTestTransitRouter() *transitRouter {
	return &transitRouter{
		name: eid.New(),
	}
}

type transitRouter struct {
	id   string
	name string
	tags map[string]interface{}
}

func (entity *transitRouter) getId() string {
	return entity.id
}

func (entity *transitRouter) setId(id string) {
	entity.id = id
}

func (entity *transitRouter) getEntityType() string {
	return "transit-routers"
}

func (entity *transitRouter) toJson(_ bool, ctx *TestContext, _ ...string) string {
	entityData := gabs.New()
	ctx.setJsonValue(entityData, entity.name, "name")
	ctx.setJsonValue(entityData, entity.tags, "tags")

	return entityData.String()
}

func (entity *transitRouter) validate(ctx *TestContext, c *gabs.Container) {
	if entity.tags == nil {
		entity.tags = map[string]interface{}{}
	}
	ctx.pathEquals(c, entity.name, path("name"))
	ctx.pathEquals(c, entity.tags, path("tags"))
}

type ca struct {
	id                        string
	name                      string                 `json:"name"`
	isAutoCaEnrollmentEnabled bool                   `json:"isAutoCaEnrollmentEnabled"`
	isAuthEnabled             bool                   `json:"isAuthEnabled"`
	isOttCaEnrollmentEnabled  bool                   `json:"isOttCaEnrollmentEnabled"`
	certPem                   string                 `json:"certPem"`
	identityRoles             []string               `json:"identityRoles"`
	identityNameFormat        string                 `json:"identityNameFormat"`
	tags                      map[string]interface{} `json:"tags"`

	privateKey crypto.Signer     `json:"-"` //utility property, not used in API calls
	publicCert *x509.Certificate `json:"-"` //utility property, not used in API calls
}

func newTestCa(identityRoles ...string) *ca {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	caCert := &x509.Certificate{
		SerialNumber: big.NewInt(2019),
		Subject: pkix.Name{
			Organization: []string{"Ziti Dev"},
			Country:      []string{"US"},
			Province:     []string{"Anywhere"},
			Locality:     []string{"Anytime"},
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(0, 0, 1),
		IsCA:                  true,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		BasicConstraintsValid: true,
	}

	caBytes, err := x509.CreateCertificate(rand.Reader, caCert, caCert, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}

	caCert, err = x509.ParseCertificate(caBytes)

	caPEM := new(bytes.Buffer)
	_ = pem.Encode(caPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: caBytes,
	})

	if identityRoles == nil {
		identityRoles = []string{}
	}

	return &ca{
		name:                      eid.New(),
		isAutoCaEnrollmentEnabled: true,
		isAuthEnabled:             true,
		isOttCaEnrollmentEnabled:  true,
		certPem:                   caPEM.String(),
		identityRoles:             identityRoles,
		identityNameFormat:        "[caName]-[commonName]-[requestedName]",
		tags:                      map[string]interface{}{},
		privateKey:                key,
		publicCert:                caCert,
	}
}

func (entity ca) getId() string {
	return entity.id
}

func (entity ca) setId(id string) {
	entity.id = id
}

func (entity ca) getEntityType() string {
	return "cas"
}

func (entity ca) toJson(create bool, ctx *TestContext, fields ...string) string {
	entityData := gabs.New()
	ctx.setValue(entityData, entity.name, fields, "name")
	ctx.setValue(entityData, entity.isOttCaEnrollmentEnabled, fields, "isOttCaEnrollmentEnabled")
	ctx.setValue(entityData, entity.isAutoCaEnrollmentEnabled, fields, "isAutoCaEnrollmentEnabled")
	ctx.setValue(entityData, entity.isAuthEnabled, fields, "isAuthEnabled")
	ctx.setValue(entityData, entity.identityRoles, fields, "identityRoles")
	ctx.setValue(entityData, entity.tags, fields, "tags")
	ctx.setValue(entityData, entity.identityNameFormat, fields, "identityNameFormat")

	if create {
		ctx.setValue(entityData, entity.certPem, fields, "certPem")
	}

	return entityData.String()
}

func (entity ca) validate(ctx *TestContext, c *gabs.Container) {
	if entity.tags == nil {
		entity.tags = map[string]interface{}{}
	}
	ctx.pathEquals(c, entity.name, path("name"))
	sort.Strings(entity.identityRoles)
	ctx.pathEqualsStringSlice(c, entity.identityRoles, path("identityRoles"))
	ctx.pathEquals(c, entity.certPem, path("certPem"))
	ctx.pathEquals(c, entity.isAuthEnabled, path("isAuthEnabled"))
	ctx.pathEquals(c, entity.isAutoCaEnrollmentEnabled, path("isAutoCaEnrollmentEnabled"))
	ctx.pathEquals(c, entity.isOttCaEnrollmentEnabled, path("isOttCaEnrollmentEnabled"))
	ctx.pathEquals(c, entity.identityNameFormat, path("identityNameFormat"))
	ctx.pathEquals(c, entity.tags, path("tags"))
}

func (entity ca) CreateSignedCert(name string) *certAuthenticator {
	clientKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	csrTemplate := &x509.CertificateRequest{
		Subject: pkix.Name{
			CommonName:   name,
			Organization: []string{"Ziti Dev"},
			Country:      []string{"US"},
			Province:     []string{"Anywhere"},
			Locality:     []string{"Anytime"},
		},
	}

	csrBytes, err := x509.CreateCertificateRequest(rand.Reader, csrTemplate, clientKey)
	if err != nil {
		panic(err)
	}

	csr, err := x509.ParseCertificateRequest(csrBytes)

	if err != nil {
		panic(err)
	}

	if err = csr.CheckSignature(); err != nil {
		panic(err)
	}

	certTemplate := x509.Certificate{
		Signature: csr.Signature,

		PublicKeyAlgorithm: csr.PublicKeyAlgorithm,
		PublicKey:          csr.PublicKey,

		SerialNumber: big.NewInt(2020),
		Issuer:       entity.publicCert.Subject,
		Subject:      csr.Subject,
		NotBefore:    time.Now().AddDate(0, 0, -1),
		NotAfter:     time.Now().AddDate(1, 0, 0),
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageDataEncipherment,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		IsCA:         false,
	}

	clientBytes, err := x509.CreateCertificate(rand.Reader, &certTemplate, entity.publicCert, csr.PublicKey, entity.privateKey)

	if err != nil {
		panic(err)
	}

	clientCert, err := x509.ParseCertificate(clientBytes)

	if err != nil {
		panic(err)
	}

	clientPEM := new(bytes.Buffer)
	_ = pem.Encode(clientPEM, &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: clientBytes,
	})

	return &certAuthenticator{
		cert:    clientCert,
		key:     clientKey,
		certPem: clientPEM.String(),
	}
}
