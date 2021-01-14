// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tls_test

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	"github.com/tigera/key-cert-provisioner/pkg/tls"
)

var _ = Describe("Test TLS functions", func() {

	Context("Test signature algorithm", func() {
		It("should translate a string to the corresponding x509 value", func() {
			// Installation Spec enum defines: "";SHA256WithRSA;SHA384WithRSA;SHA512WithRSA;ECDSAWithSHA256;ECDSAWithSHA384;ECDSAWithSHA512;
			Expect(tls.SignatureAlgorithm("")).To(Equal(x509.SHA256WithRSA))
			Expect(tls.SignatureAlgorithm("SHA256WithRSA")).To(Equal(x509.SHA256WithRSA))
			Expect(tls.SignatureAlgorithm("SHA384WithRSA")).To(Equal(x509.SHA384WithRSA))
			Expect(tls.SignatureAlgorithm("SHA512WithRSA")).To(Equal(x509.SHA512WithRSA))
			Expect(tls.SignatureAlgorithm("ECDSAWithSHA256")).To(Equal(x509.ECDSAWithSHA256))
			Expect(tls.SignatureAlgorithm("ECDSAWithSHA384")).To(Equal(x509.ECDSAWithSHA384))
			Expect(tls.SignatureAlgorithm("ECDSAWithSHA512")).To(Equal(x509.ECDSAWithSHA512))
		})
	})

	DescribeTable("Test private key algorithm", func(keyAlg string, expectRSA bool, expectErr bool) {
		key, pem, err := tls.GeneratePrivateKey(keyAlg)
		if expectErr {
			Expect(err).To(HaveOccurred())
		} else {
			Expect(err).NotTo(HaveOccurred())
			Expect(pem).ToNot(HaveLen(0))

			if expectRSA {
				_, isRsa := key.(*rsa.PrivateKey)
				Expect(isRsa).To(BeTrue())
			} else {
				_, isECDSA := key.(*ecdsa.PrivateKey)
				Expect(isECDSA).To(BeTrue())
			}
		}
		Expect(err != nil).To(Equal(expectErr))
	},
		// Installation Spec enum defines: "";RSAWithSize2048;RSAWithSize4096;RSAWithSize8192;ECDSAWithCurve256;ECDSAWithCurve384;ECDSAWithCurve521;
		Entry("Empty -> rsa private key", "", true, false),
		Entry("RSAWithSize2048 -> rsa private key", "RSAWithSize2048", true, false),
		Entry("RSAWithSize4096 -> rsa private key ", "RSAWithSize4096", true, false),
		Entry("RSAWithSize8192 -> rsa private key", "RSAWithSize8192", true, false),
		Entry("ECDSAWithCurve256 -> ecdsa private key", "ECDSAWithCurve256", false, false),
		Entry("ECDSAWithCurve384 -> ecdsa private key", "ECDSAWithCurve384", false, false),
		Entry("ECDSAWithCurve521 -> ecdsa private key", "ECDSAWithCurve521", false, false),
	)

})
