package cognitoAccess

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

var cognitoClient *cognitoidentityprovider.Client

func init() {
	// Load the AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create the Cognito Identity Provider client
	cognitoClient = cognitoidentityprovider.NewFromConfig(cfg)
}

func CreateUser(email string, password string, userAtrributes []types.AttributeType, ctx context.Context, userPoolId string) (*cognitoidentityprovider.AdminCreateUserOutput, error) {
	//create user cognito
	input := &cognitoidentityprovider.AdminCreateUserInput{
		// DesiredDeliveryMediums: []*string{
		//     aws.String("SMS"),
		// },
		// MessageAction: aws.String("SUPPRESS"),
		TemporaryPassword: aws.String(password),
		UserAttributes:    userAtrributes,
		UserPoolId:        aws.String(userPoolId),
		Username:          aws.String(email),
	}

	result, err := cognitoClient.AdminCreateUser(ctx, input)
	if err != nil {
		log.Println("error create cognito user", err)
		return result, err
	}

	//account confirmation
	input1 := &cognitoidentityprovider.AdminSetUserPasswordInput{
		UserPoolId: aws.String(userPoolId),
		Username:   aws.String(email),
		Password:   aws.String(password),
		Permanent:  true,
	}
	//"ap-southeast-1_OtBDMeQb3"
	result1, err := cognitoClient.AdminSetUserPassword(ctx, input1)
	log.Println("result cognito", result1)
	if err != nil {
		log.Println("error confirm cognito user", err)
		return result, err
	}

	return result, err
}

func UpdateUser(username string, userAttributes []types.AttributeType, ctx context.Context, userPoolId string) (*cognitoidentityprovider.AdminUpdateUserAttributesOutput, error) {
	//create user cognito
	input := &cognitoidentityprovider.AdminUpdateUserAttributesInput{
		UserAttributes: userAttributes,
		UserPoolId:     aws.String(userPoolId),
		Username:       aws.String(username),
	}

	result, err := cognitoClient.AdminUpdateUserAttributes(ctx, input)
	if err != nil {
		log.Println("error update cognito user", err)
		return result, err
	}

	return result, err
}

func EditPassword(ctx context.Context, userPoolId string, username string, password string) (*cognitoidentityprovider.AdminSetUserPasswordOutput, error) {
	//account confirmation
	input1 := &cognitoidentityprovider.AdminSetUserPasswordInput{
		UserPoolId: aws.String(userPoolId),
		Username:   aws.String(username),
		Password:   aws.String(password),
		Permanent:  true,
	}

	result, err := cognitoClient.AdminSetUserPassword(ctx, input1)
	log.Println("result cognito", result)
	if err != nil {
		log.Println("error confirm cognito user", err)
		return result, err
	}

	return result, err
}

func DeleteUser(username string, ctx context.Context, userPoolId string) (*cognitoidentityprovider.AdminDeleteUserOutput, error) {
	//create user cognito
	input := &cognitoidentityprovider.AdminDeleteUserInput{
		UserPoolId: aws.String(userPoolId),
		Username:   aws.String(username),
	}

	result, err := cognitoClient.AdminDeleteUser(ctx, input)
	if err != nil {
		log.Println("error delete cognito user", err)
		return result, err
	}

	return result, err
}

func GetUser(username string, ctx context.Context, userPoolId string) (*cognitoidentityprovider.AdminGetUserOutput, error) {
	//create user cognito
	input := &cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: aws.String(userPoolId),
		Username:   aws.String(username),
	}

	result, err := cognitoClient.AdminGetUser(ctx, input)
	if err != nil {
		log.Println("error get cognito user", err)
		return result, err
	}

	return result, err
}

func ListUsers(ctx context.Context, userPoolId string, attributesToGet []string) (*cognitoidentityprovider.ListUsersOutput, error) {
	//create user cognito
	input := &cognitoidentityprovider.ListUsersInput{
		UserPoolId:      aws.String(userPoolId),
		AttributesToGet: attributesToGet,
		// Filter:          aws.String(filter),
	}

	result, err := cognitoClient.ListUsers(ctx, input)
	if err != nil {
		log.Println("error list cognito user", err)
		return result, err
	}

	return result, err
}
