package plugin

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Default registry", func() {
	BeforeEach(func() {
		Register(testPluginType(), testPluginName, newTestPlugin)
	})
	AfterEach(func() {
		defaultRegistry = newTypeRegistry()
	})
	It("lookup", func() {
		Expect(Lookup(testPluginType())).To(BeTrue())
	})
	It("lookup factory", func() {
		Expect(LookupFactory(testPluginFactoryType())).To(BeTrue())
	})
	It("new", func() {
		plugin, err := New(testPluginType(), testPluginName)
		Expect(err).NotTo(HaveOccurred())
		Expect(plugin).NotTo(BeNil())
	})
	It("new factory", func() {
		pluginFactory, err := NewFactory(testPluginFactoryType(), testPluginName)
		Expect(err).NotTo(HaveOccurred())
		Expect(pluginFactory).NotTo(BeNil())
	})
})

var _ = Describe("type helpers", func() {
	It("ptr type", func() {
		var plugin testPluginInterface
		Expect(PtrType(&plugin)).To(Equal(testPluginType()))
	})
	It("factory plugin type ok", func() {
		factoryPlugin, ok := FactoryPluginType(testPluginFactoryType())
		Expect(ok).To(BeTrue())
		Expect(factoryPlugin).To(Equal(testPluginType()))
	})
})
