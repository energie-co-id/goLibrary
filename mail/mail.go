package mail

import (
	"log"

	"gopkg.in/mail.v2"
)

func SendMail(subject string, emailReceiver string, body string) error {
	// Set up the email message
	m := mail.NewMessage()
	m.SetHeader("From", "evgate-support@stroomer.id")
	m.SetHeader("To", emailReceiver)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	// Set up the email server
	d := mail.NewDialer("smtp.gmail.com", 587, "jauhar@harapanenergie.com", "jdwxqxscgznopekl")

	// Send the email
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	log.Println("Email sent successfully!")
	return nil
}

func StopCharging() string {
	mail := `<!DOCTYPE html>
    <html>
    <head>
        <title>EV Gate Notification</title>
        <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css" integrity="sha384-..." crossorigin="anonymous">
    </head>
    <body style="font-family: Arial, sans-serif; line-height: 1.6; background-color: #f2f2f2; margin: 0; padding: 0;">
        <table width="100%" border="0" cellspacing="0" cellpadding="0">
            <tr>
                <td align="center">
                    <table width="600" border="0" cellspacing="0" cellpadding="0" style="background-color: #ffffff;">
                        <tr>
                            <td align="center" style="padding: 10px;">
                                <img src="https://evgate.energie.co.id/assets/images/logo/swa-logo.png" alt="Stroomer EV Gate" style="max-width: 200px;">
                            </td>
                        </tr>
                        <tr>
                            <td style="padding: 20px; margin-bottom: 0;">
                                <h4>Hello</h4>
                                    <p>Your charger <b>{{name}}</b> has been stopped, thank you for using our products</p>
                                <h4>Details : </h4>
                                <hr style="border:1px solid rgb(235, 231, 231);"/>
                            </td>
                        </tr>
                        <tr>
                            <td style="padding: 20px;">
                                <table width="100%" border="0" cellspacing="0" cellpadding="8">
                                    <tr align="left">
                                        <th style="border-bottom: 1px solid #ddd;">Transaction ID</th>
                                        <td style="border-bottom: 1px solid #ddd;">:  {{id}}</td>
                                    </tr>
                                    <tr align="left">
                                        <th style="border-bottom: 1px solid #ddd;">Car Type</th>
                                        <td style="border-bottom: 1px solid #ddd;">:  {{car_type}}</td>
                                    </tr>
                                    <tr align="left">
                                        <th style="border-bottom: 1px solid #ddd;">Energy</th>
                                        <td style="border-bottom: 1px solid #ddd;">:  {{energy_html}}</td>
                                    </tr>
                                    <tr align="left">
                                        <th style="border-bottom: 1px solid #ddd;">Duration</th>
                                        <td style="border-bottom: 1px solid #ddd;">:  {{duration}}</td>
                                    </tr>
                                    <tr align="left">
                                        <th style="border-bottom: 1px solid #ddd;">Start Time</th>
                                        <td style="border-bottom: 1px solid #ddd;">:  {{start_time}}</td>
                                    </tr>
                                    <tr align="left">
                                        <th style="border-bottom: 1px solid #ddd;">Stop Time</th>
                                        <td style="border-bottom: 1px solid #ddd;">:  {{stop_time}}</td>
                                    </tr>
                                </table>
                                <tr>
                                    <td style="padding-left: 20px; margin-bottom: 0;"><b>Issues with this notification?</b></td>
                                </tr>
                                <tr>
                                    <td style="padding-bottom: 20px; padding-left: 20px; margin-bottom: 0;">You can reply to this email</td>
                                </tr>
                                
                            </td>
                        </tr>
                        <tr>
                            <td align="center" style="background-color: #f4a106; padding: 5px; color: #ffffff;">
                                <p>Stroomer Energie 2023</p>
                            </td>
                        </tr>
                    </table>
                </td>
            </tr>
        </table>
    </body>
    </html>`

	return mail
}
func TopupSuccess() string {
	mail := `<!DOCTYPE html>
      <html>
      <head>
          <title>EV Gate Notification</title>
          <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css" integrity="sha384-..." crossorigin="anonymous">
      </head>
      <body style="font-family: Arial, sans-serif; line-height: 1.6; background-color: #f2f2f2; margin: 0; padding: 0;">
          <table width="100%" border="0" cellspacing="0" cellpadding="0">
              <tr>
                  <td align="center">
                      <table width="600" border="0" cellspacing="0" cellpadding="0" style="background-color: #ffffff;">
                          <tr>
                              <td align="center" style="padding: 10px;">
                                  <img src="https://evgate.energie.co.id/assets/images/logo/swa-logo.png" alt="Stroomer EV Gate" style="max-width: 200px;">
                              </td>
                          </tr>
                          <tr>
                              <td style="padding: 20px; margin-bottom: 0;">
                                  <h1 style="color: #4fd971;">Top Up Success</h1>
                                  <h4>Hello</h4>
                                      <p> Your prepaid charger meter is almost empty, and has been refilled with {{energy}} kWh, thank you for using our products</p>
                                  <h4>Details : </h4>
                                  <hr style="border:1px solid rgb(235, 231, 231);"/>
                              </td>
                          </tr>
                          <tr>
                              <td style="padding: 20px;">
                                  <table width="100%" border="0" cellspacing="0" cellpadding="8">
                                      <tr align="left">
                                          <th style="border-bottom: 1px solid #ddd;">EV Gate</th>
                                          <td style="border-bottom: 1px solid #ddd;">:  {{evgate}}</td>
                                      </tr>
                                      <tr align="left">
                                          <th style="border-bottom: 1px solid #ddd;">Meter ID</th>
                                          <td style="border-bottom: 1px solid #ddd;">:  {{meter_id}}</td>
                                      </tr>
                                      <tr align="left">
                                          <th style="border-bottom: 1px solid #ddd;">Time</th>
                                          <td style="border-bottom: 1px solid #ddd;">:  {{time}}</td>
                                      </tr>
                                      <tr align="left">
                                          <th style="border-bottom: 1px solid #ddd;">Token</th>
                                          <td style="border-bottom: 1px solid #ddd;">:  {{token}}</td>
                                      </tr>
                                      <tr align="left">
                                          <th style="border-bottom: 1px solid #ddd;">Energy</th>
                                          <td style="border-bottom: 1px solid #ddd;">:  {{energy}} kWh</td>
                                      </tr>
                                      <tr align="left">
                                          <th style="border-bottom: 1px solid #ddd;">Credit</th>
                                          <td style="border-bottom: 1px solid #ddd;">:  {{last_credit}} kWh</td>
                                      </tr>
                                  </table>
                                  <tr>
                                      <td style="padding-left: 20px; margin-bottom: 0;"><b>Issues with this notification?</b></td>
                                  </tr>
                                  <tr>
                                      <td style="padding-bottom: 20px; padding-left: 20px; margin-bottom: 0;">You can reply to this email</td>
                                  </tr>
                                  
                              </td>
                          </tr>
                          <tr>
                              <td align="center" style="background-color: #f4a106; padding: 5px; color: #ffffff;">
                                  <p>Stroomer Energie 2023</p>
                              </td>
                          </tr>
                      </table>
                  </td>
              </tr>
          </table>
      </body>
      </html>`

	return mail

}

func TopupFailed() string {
	mail := `<!DOCTYPE html>
      <html>
      <head>
          <title>EV Gate Notification</title>
          <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.0.0-beta3/css/all.min.css" integrity="sha384-..." crossorigin="anonymous">
      </head>
      <body style="font-family: Arial, sans-serif; line-height: 1.6; background-color: #f2f2f2; margin: 0; padding: 0;">
          <table width="100%" border="0" cellspacing="0" cellpadding="0">
              <tr>
                  <td align="center">
                      <table width="600" border="0" cellspacing="0" cellpadding="0" style="background-color: #ffffff;">
                          <tr>
                              <td align="center" style="padding: 10px;">
                                  <img src="https://evgate.energie.co.id/assets/images/logo/swa-logo.png" alt="Stroomer EV Gate" style="max-width: 200px;">
                              </td>
                          </tr>
                          <tr>
                              <td style="padding: 20px; margin-bottom: 0;">
                                  <h1 style="color: #d9534f;">Top Up Failed</h1>
                                  <h4>Hello</h4>
                                      <p> Your prepaid charger meter is almost empty, but we detected an error when top up the prepaid charger meter</p>
                                      <p style="color: #d9534f;"> Please double check the token or meter ID that you entered</p>
                                  <h4>Details : </h4>
                                  <hr style="border:1px solid rgb(235, 231, 231);"/>
                              </td>
                          </tr>
                          <tr>
                              <td style="padding: 20px;">
                                  <table width="100%" border="0" cellspacing="0" cellpadding="8">
                                      <tr align="left">
                                          <th style="border-bottom: 1px solid #ddd;">EV Gate</th>
                                          <td style="border-bottom: 1px solid #ddd;">:  {{evgate}}</td>
                                      </tr>
                                      <tr align="left">
                                          <th style="border-bottom: 1px solid #ddd;">Meter ID</th>
                                          <td style="border-bottom: 1px solid #ddd;">:  {{meter_id}}</td>
                                      </tr>
                                      <tr align="left">
                                          <th style="border-bottom: 1px solid #ddd;">Time</th>
                                          <td style="border-bottom: 1px solid #ddd;">:  {{time}}</td>
                                      </tr>
                                      <tr align="left">
                                          <th style="border-bottom: 1px solid #ddd;">Token</th>
                                          <td style="border-bottom: 1px solid #ddd;">:  {{token}}</td>
                                      </tr>
                                      <tr align="left">
                                          <th style="border-bottom: 1px solid #ddd;">Energy</th>
                                          <td style="border-bottom: 1px solid #ddd; color: #d9534f">:  {{energy}} kWh</td>
                                      </tr>
                                      <tr align="left">
                                          <th style="border-bottom: 1px solid #ddd;">Credit</th>
                                          <td style="border-bottom: 1px solid #ddd;">:  {{last_credit}} kWh</td>
                                      </tr>
                                  </table>
                                  <tr>
                                      <td style="padding-left: 20px; margin-bottom: 0;"><b>Issues with this notification?</b></td>
                                  </tr>
                                  <tr>
                                      <td style="padding-bottom: 20px; padding-left: 20px; margin-bottom: 0;">You can reply to this email</td>
                                  </tr>
                                  
                              </td>
                          </tr>
                          <tr>
                              <td align="center" style="background-color: #d9534f; padding: 5px; color: #ffffff;">
                                  <p>Stroomer Energie 2023</p>
                              </td>
                          </tr>
                      </table>
                  </td>
              </tr>
          </table>
      </body>
      </html>`

	return mail
}
