response.StatusCode = 404
response.ContentType = "json"
response.Body = JSON.stringify({
    message: "user not found"
})