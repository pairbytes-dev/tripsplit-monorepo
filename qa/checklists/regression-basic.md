# TripSplit – Simple Regression Checklist

## Registration (Sign up)

- [ ] User can register with valid name, email and password.  
- [ ] Validation errors appear for missing required fields (name, email, password).  
- [ ] Validation error appears for invalid email format.  
- [ ] Validation error appears for weak/too short password.  
- [ ] Success feedback is shown after successful registration.

## Login

- [ ] User can log in with valid email and password.  
- [ ] Login fails with clear error message when password is incorrect.  
- [ ] Login fails with clear error message when email is not registered.  
- [ ] After successful login, user is redirected to the authenticated home screen.

## Authenticated home screen

- [ ] After login, the authenticated home screen is displayed.  
- [ ] If the user has trips, the list of trips is shown correctly.  
- [ ] If the user has no trips, an empty state is displayed with a call-to-action to create a trip.  
- [ ] User can navigate from a trip in the list to its detail screen (if already implemented).

## Access control / security

- [ ] When not authenticated and trying to access the authenticated home URL directly, the user is redirected to the login page.  
- [ ] After logout, the user cannot access authenticated pages via browser back button (must be redirected to login).  
- [ ] Session/token handling works correctly across page reloads (user stays logged in while session is valid).
