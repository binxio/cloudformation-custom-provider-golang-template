package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type resourceProperties struct {
	Value string
}

func validateInput(event cfn.Event) (result *resourceProperties, err error) {
	result = new(resourceProperties)
	if value, ok := event.ResourceProperties["Value"].(string); ok {
		result.Value = value
	} else {
		return nil, fmt.Errorf("Value not specified or not a string")
	}
	return result, nil
}

func create(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	var props *resourceProperties

	if props, err = validateInput(event); err != nil {
		return "", nil, err
	}

	data = make(map[string]interface{})
	data["Value"] = props.Value
	return props.Value, data, nil
}

func update(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	var props *resourceProperties

	if props, err = validateInput(event); err != nil {
		return "", nil, err
	}

	data = make(map[string]interface{})
	data["Value"] = props.Value
	return props.Value, data, nil
}

func delete(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	if physicalResourceID != "" {
		log.Printf("Deleting %s", physicalResourceID)
	} else {
		log.Printf("ignoring invalid physical resource id")
	}
	return "", nil, nil
}

func handler(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {

	if event.ResourceType == "Custom::ContainerImage" {
		switch event.RequestType {
		case cfn.RequestCreate:
			return create(ctx, event)

		case cfn.RequestUpdate:
			return update(ctx, event)

		case cfn.RequestDelete:
			return delete(ctx, event)
		default:
			return "", nil, fmt.Errorf("unsupported request type: %s", event.RequestType)
		}
	}
	return "", nil, fmt.Errorf("unsupported resource type: %s", event.ResourceType)
}

func main() {
	lambda.Start(cfn.LambdaWrap(handler))
}
