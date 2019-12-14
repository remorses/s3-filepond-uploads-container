# Upload you file to S3 using filepond and a docker container

Container that receives uploads at path `/upload` and send the files in an s3 bucket.
The files Key and url are made with a random string so that these files can be exposed with `public-read` ACL and used directly in your app,
Supports uploading with multipart, can upload only one file at a time, make multiple requests to upload more files
You can also test that uploading works via curl: 
`curl -X POST -F "file=@path/to/file.png" http://this-container/upload`

Can be used with [`filepond`](https://pqina.nl/filepond/) using the container url as the `server` parameter.


The image is 20 Mb uncompressed, the files are handled as streams so memory usage shoud not be high

To see an example add the required env vars `ACCESS_KEY_ID` `SECRET_ACCESS_KEY` in an `.env` file and run in different terminal tabs
`docker compose up --build`

`cd client && yarn && yarn dev`

# How to use

Add the service to your docker-compose
Then your client can then use `fiepond-react` with the container url `http://localhost:8010/upload`
The service returns the url of the uploaded file, automatically handled by filepond that set it to the File object


```yml
version: '3'
services:
    upload:
        image: xmorse/s3-filepond
        ports:
            - 8010:80
        environment:
            - AWS_ACCESS_KEY_ID=$ACCESS_KEY_ID
            - AWS_SECRET_ACCESS_KEY=$SECRET_ACCESS_KEY
            - DIRECTORY=testing-filepond/
            - BUCKET=testshitkjhdgfslkjsdbf
            - REGION=eu-west-1
```

```tsx
import React, { useState } from 'react'
import { FilePond, File } from 'react-filepond'
import 'filepond/dist/filepond.min.css'

const SERVER = 'http://localhost:8010/upload' // this container url

export default function App() {
    const [files, setFiles] = useState<File[]>([])
    const [urls, setUrls] = useState<string[]>([])
    return (
        <div className='App'>
            <FilePond
                files={files}
                allowMultiple={true}
                onupdatefiles={setFiles}
                onprocessfile={(err, file) => {
                    setUrls([...urls, file.serverId])
                }}
                labelIdle='Drag & Drop your files or <span class="filepond--label-action">Browse</span>'
                server={SERVER}
            />
            <pre>{JSON.stringify(urls, null, 4)}</pre>
        </div>
    )
}
```
