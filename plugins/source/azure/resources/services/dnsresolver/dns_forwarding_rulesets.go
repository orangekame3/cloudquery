// Code generated by codegen; DO NOT EDIT.
package dnsresolver

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dnsresolver/armdnsresolver"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func DnsForwardingRulesets() *schema.Table {
	return &schema.Table{
		Name:      "azure_dnsresolver_dns_forwarding_rulesets",
		Resolver:  fetchDnsForwardingRulesets,
		Multiplex: client.SubscriptionMultiplexRegisteredNamespace("azure_dnsresolver_dns_forwarding_rulesets", client.Namespacemicrosoft_network),
		Transform: transformers.TransformWithStruct(&armdnsresolver.DNSForwardingRuleset{}),
		Columns: []schema.Column{
			{
				Name:     "subscription_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAzureSubscription,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchDnsForwardingRulesets(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc, err := armdnsresolver.NewDNSForwardingRulesetsClient(cl.SubscriptionId, cl.Creds, cl.Options)
	if err != nil {
		return err
	}
	pager := svc.NewListPager(nil)
	for pager.More() {
		p, err := pager.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- p.Value
	}
	return nil
}
