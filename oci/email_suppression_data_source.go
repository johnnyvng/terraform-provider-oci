// Copyright (c) 2017, 2019, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"context"

	"github.com/hashicorp/terraform/helper/schema"
	oci_email "github.com/oracle/oci-go-sdk/email"
)

func EmailSuppressionDataSource() *schema.Resource {
	return &schema.Resource{
		Read: readSingularEmailSuppression,
		Schema: map[string]*schema.Schema{
			"suppression_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			// Computed
			"compartment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"reason": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"time_created": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func readSingularEmailSuppression(d *schema.ResourceData, m interface{}) error {
	sync := &EmailSuppressionDataSourceCrud{}
	sync.D = d
	sync.Client = m.(*OracleClients).emailClient

	return ReadResource(sync)
}

type EmailSuppressionDataSourceCrud struct {
	D      *schema.ResourceData
	Client *oci_email.EmailClient
	Res    *oci_email.GetSuppressionResponse
}

func (s *EmailSuppressionDataSourceCrud) VoidState() {
	s.D.SetId("")
}

func (s *EmailSuppressionDataSourceCrud) Get() error {
	request := oci_email.GetSuppressionRequest{}

	if suppressionId, ok := s.D.GetOkExists("suppression_id"); ok {
		tmp := suppressionId.(string)
		request.SuppressionId = &tmp
	}

	request.RequestMetadata.RetryPolicy = getRetryPolicy(false, "email")

	response, err := s.Client.GetSuppression(context.Background(), request)
	if err != nil {
		return err
	}

	s.Res = &response
	return nil
}

func (s *EmailSuppressionDataSourceCrud) SetData() error {
	if s.Res == nil {
		return nil
	}

	s.D.SetId(*s.Res.Id)

	if s.Res.CompartmentId != nil {
		s.D.Set("compartment_id", *s.Res.CompartmentId)
	}

	if s.Res.EmailAddress != nil {
		s.D.Set("email_address", *s.Res.EmailAddress)
	}

	s.D.Set("reason", s.Res.Reason)

	if s.Res.TimeCreated != nil {
		s.D.Set("time_created", s.Res.TimeCreated.String())
	}

	return nil
}
