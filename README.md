# ms-file

# routes
## Upload file
``curl --location 'SERVER_URL:3000/files' \
--form 'file=@"FILE_PATH.jpg"'``

## Get file
``curl --location 'SERVER_URL:3000/files/FILE_ID'``
