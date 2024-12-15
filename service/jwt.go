package service

import (
	"ecommerce-rbac/config"
	"ecommerce-rbac/logger"
	"ecommerce-rbac/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	// "gorm.io/gorm/logger"
)

func GenerateJwtToken(user models.User) (string, error) {
	expirationTime := time.Now().Add(5 * time.Hour) // Set expiration time for the token

	claims := models.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime), // Set expiration claim
		},
		Username: user.Username,
		Role:     user.Role, // Assuming Role is a slice of constants.Role
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)                       // Create a new token with claims
	tokenString, err := token.SignedString([]byte(config.AppConfig.JwtSecretString)) // Sign the token with the secret
	if err != nil {

		logger.Error(err)
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(jwtToken string, requiredRole string) (string, error) {
	logger.Debug("JWT Token: ", jwtToken)
	logger.Debug("Required Role: ", requiredRole)

	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.AppConfig.JwtSecretString), nil
	})
	if err != nil {
		logger.Error("Validation failed: ", err)
		return "", errors.New("validation failed")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		logger.Debug("Claims: ", claims) // Log claims for debugging

		// Check if the role exists in claims
		if roles, ok := claims["role"].([]interface{}); ok {
			for _, value := range roles {
				if requiredRole == value.(string) {
					username, ok := claims["username"].(string)
					if !ok {
						logger.Error("Error retrieving Username from claims")
						return "", errors.New("error retrieving Username from claims")
					}
					return username, nil 
				}
			}
			logger.Error("Required role not found in claims")
			return "", errors.New("authorization mismatch")
		} else if role, ok := claims["role"].(string); ok { // Handle single role case
			if requiredRole == role {
				username, ok := claims["username"].(string)
				if !ok {
					logger.Error("Error retrieving Username from claims")
					return "", errors.New("error retrieving Username from claims")
				}
				return username, nil 
			}
			logger.Error("Required role not found in claims")
			return "", errors.New("authorization mismatch")
		} else {
			logger.Error("Unable to authorize: role not found in claims")
			return "", errors.New("authorization mismatch")
		}
	}

	return "", errors.New("authorization mismatch")
}



func ValidateTokenWithMultipleRole(jwtToken string, requiredRoles []string) (models.UserClaims, error) {
    logger.Debug("JWT Token: ", jwtToken)
    logger.Debug("Required Roles: ", requiredRoles)

    token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
        // Validate the signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, errors.New("unexpected signing method")
        }
        return []byte(config.AppConfig.JwtSecretString), nil
    })
    if err != nil {
        logger.Error("Validation failed: ", err)
        return models.UserClaims{}, errors.New("validation failed")
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        logger.Debug("Claims: ", claims) // Log claims for debugging

        // Check if the role exists in claims
        if userRoles, ok := claims["role"].([]interface{}); ok {
            for _, value := range userRoles {
                for _, requiredRole := range requiredRoles { // Check against all required roles
                    if requiredRole == value.(string) {
                        username, ok := claims["username"].(string)
                        if !ok {
                            logger.Error("Error retrieving Username from claims")
                            return models.UserClaims{}, errors.New("error retrieving Username from claims")
                        }
                        return models.UserClaims{Username: username, Role: requiredRole}, nil // Return both username and role
                    }
                }
            }
        } else if role, ok := claims["role"].(string); ok { // Handle single role case
            for _, requiredRole := range requiredRoles {
                if requiredRole == role {
                    username, ok := claims["username"].(string)
                    if !ok {
                        logger.Error("Error retrieving Username from claims")
                        return models.UserClaims{}, errors.New("error retrieving Username from claims")
                    }
                    return models.UserClaims{Username: username, Role: role}, nil // Return both username and role
                }
            }
        } else {
            logger.Error("Unable to authorize: role not found in claims")
            return models.UserClaims{}, errors.New("authorization mismatch")
        }
    }

    return models.UserClaims{}, errors.New("authorization mismatch")
}

