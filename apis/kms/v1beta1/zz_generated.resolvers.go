// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by angryjet. DO NOT EDIT.
// Code transformed by upjet. DO NOT EDIT.

package v1beta1

import (
	"context"
	reference "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	resource "github.com/crossplane/upjet/pkg/resource"
	errors "github.com/pkg/errors"
	common "github.com/upbound/provider-gcp/config/common"
	apisresolver "github.com/upbound/provider-gcp/internal/apis"
	client "sigs.k8s.io/controller-runtime/pkg/client"
)

func (mg *CryptoKey) ResolveReferences( // ResolveReferences of this CryptoKey.
	ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta1", "KeyRing", "KeyRingList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyRing),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.KeyRingRef,
			Selector:     mg.Spec.ForProvider.KeyRingSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyRing")
	}
	mg.Spec.ForProvider.KeyRing = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyRingRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this CryptoKeyIAMMember.
func (mg *CryptoKeyIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta1", "CryptoKey", "CryptoKeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CryptoKeyID),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.CryptoKeyIDRef,
			Selector:     mg.Spec.ForProvider.CryptoKeyIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CryptoKeyID")
	}
	mg.Spec.ForProvider.CryptoKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CryptoKeyIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta1", "CryptoKey", "CryptoKeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CryptoKeyID),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.CryptoKeyIDRef,
			Selector:     mg.Spec.InitProvider.CryptoKeyIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CryptoKeyID")
	}
	mg.Spec.InitProvider.CryptoKeyID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CryptoKeyIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this CryptoKeyVersion.
func (mg *CryptoKeyVersion) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta1", "CryptoKey", "CryptoKeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CryptoKey),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.CryptoKeyRef,
			Selector:     mg.Spec.ForProvider.CryptoKeySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CryptoKey")
	}
	mg.Spec.ForProvider.CryptoKey = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CryptoKeyRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta1", "CryptoKey", "CryptoKeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CryptoKey),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.CryptoKeyRef,
			Selector:     mg.Spec.InitProvider.CryptoKeySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CryptoKey")
	}
	mg.Spec.InitProvider.CryptoKey = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CryptoKeyRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this KeyRingIAMMember.
func (mg *KeyRingIAMMember) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta1", "KeyRing", "KeyRingList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyRingID),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.KeyRingIDRef,
			Selector:     mg.Spec.ForProvider.KeyRingIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyRingID")
	}
	mg.Spec.ForProvider.KeyRingID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyRingIDRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta1", "KeyRing", "KeyRingList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.KeyRingID),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.KeyRingIDRef,
			Selector:     mg.Spec.InitProvider.KeyRingIDSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.KeyRingID")
	}
	mg.Spec.InitProvider.KeyRingID = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.KeyRingIDRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this KeyRingImportJob.
func (mg *KeyRingImportJob) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta1", "KeyRing", "KeyRingList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.KeyRing),
			Extract:      common.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.KeyRingRef,
			Selector:     mg.Spec.ForProvider.KeyRingSelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.KeyRing")
	}
	mg.Spec.ForProvider.KeyRing = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.KeyRingRef = rsp.ResolvedReference

	return nil
}

// ResolveReferences of this SecretCiphertext.
func (mg *SecretCiphertext) ResolveReferences(ctx context.Context, c client.Reader) error {
	var m xpresource.Managed
	var l xpresource.ManagedList
	r := reference.NewAPIResolver(c, mg)

	var rsp reference.ResolutionResponse
	var err error
	{
		m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta1", "CryptoKey", "CryptoKeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.ForProvider.CryptoKey),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.ForProvider.CryptoKeyRef,
			Selector:     mg.Spec.ForProvider.CryptoKeySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.ForProvider.CryptoKey")
	}
	mg.Spec.ForProvider.CryptoKey = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.ForProvider.CryptoKeyRef = rsp.ResolvedReference
	{
		m, l, err = apisresolver.GetManagedResource("kms.gcp.upbound.io", "v1beta1", "CryptoKey", "CryptoKeyList")
		if err != nil {
			return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
		}

		rsp, err = r.Resolve(ctx, reference.ResolutionRequest{
			CurrentValue: reference.FromPtrValue(mg.Spec.InitProvider.CryptoKey),
			Extract:      resource.ExtractResourceID(),
			Reference:    mg.Spec.InitProvider.CryptoKeyRef,
			Selector:     mg.Spec.InitProvider.CryptoKeySelector,
			To:           reference.To{List: l, Managed: m},
		})
	}
	if err != nil {
		return errors.Wrap(err, "mg.Spec.InitProvider.CryptoKey")
	}
	mg.Spec.InitProvider.CryptoKey = reference.ToPtrValue(rsp.ResolvedValue)
	mg.Spec.InitProvider.CryptoKeyRef = rsp.ResolvedReference

	return nil
}
