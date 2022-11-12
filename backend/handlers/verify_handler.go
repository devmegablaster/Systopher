package handlers

import (
	"fmt"

	"github.com/axrav/Systopher/backend/helpers"
	"github.com/axrav/Systopher/backend/types"
	"github.com/gofiber/fiber/v2"
)

func Verify(c *fiber.Ctx) error {
	resp := new(types.OTPResponse)
	if err := c.BodyParser(resp); err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Wrong data",
		})
	} else {
		if resp.Otp == "" || resp.Email == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Missing email or otp",
			})
		}
		out := helpers.GetVerified(resp.Email)
		if out {
			return c.Status(409).JSON(fiber.Map{
				"message": "the user is already verified",
			})
		}
		verified := helpers.VerifyOtp(resp.Email, resp.Otp)
		if verified {
			if err := helpers.SetVerify(resp.Email); err != nil {
				fmt.Println(err)
				return c.Status(500).JSON(fiber.Map{
					"message": "Internal server error",
				})
			}
			token, err := helpers.GenerateJWT(resp.Email, false)
			if err != nil {
				fmt.Println(err)
				return c.Status(500).JSON(fiber.Map{
					"message": "Internal server error",
				})
			}
			return c.JSON(fiber.Map{
				"message": "Verified",
				"token":   token,
			})
		} else {
			return c.Status(400).JSON(fiber.Map{
				"message": "Wrong otp",
			})
		}
	}
}

func ResendOTP(c *fiber.Ctx) error {
	email := new(types.Email)
	if err := c.BodyParser(email); err != nil {
		fmt.Println(err)
		return c.Status(500).JSON(fiber.Map{
			"message": "Wrong data",
		})
	} else {
		if email.Email == "" {
			return c.Status(400).JSON(fiber.Map{
				"message": "Missing email",
			})
		}
		out := helpers.GetVerified(email.Email)
		if out {
			return c.Status(409).JSON(fiber.Map{
				"message": "the user is already verified",
			})
		}

		sent, err := helpers.SendOtp(email.Email)
		if err != nil {
			fmt.Println(err)
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
		output := helpers.SaveOtp(email.Email, sent)

		if output {
			return c.JSON(fiber.Map{
				"message": "OTP sent",
			})
		}
	}
	return c.Status(500).JSON(fiber.Map{
		"message": "Internal server error",
	})
}
