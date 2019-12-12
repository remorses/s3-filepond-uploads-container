# Upload you file to S3 using filepond and a docker container

To see an example add the required env vars `ACCESS_KEY_ID` `SECRET_ACCESS_KEY` in an `.env` file and run in different terminal tabs
`docker compose up --build`
`cd client && yarn && yarn dev`

# How to use
```yml
version: '3'
services:
    upload:
        build: .
        ports:
            - 8010:80
        environment:
            - ACCESS_KEY_ID=$ACCESS_KEY_ID
            - SECRET_ACCESS_KEY=$SECRET_ACCESS_KEY
            - DIRECTORY=testing-filepond/ # optional
            - BUCKET=testshitkjhdgfslkjsdbf
            - REGION=eu-west-1
```

Then your client can then use `fiepond-react` like this

```tsx
import React, { useState } from 'react'
import { FilePond, File } from 'react-filepond'
import 'filepond/dist/filepond.min.css'

const SERVER = 'http://localhost:8010/upload' // this container url

export default function App() {
    const [files, setFiles] = useState<File[]>([])
    return (
        <div className='App'>
            <FilePond
                files={files}
                allowMultiple={true}
                onupdatefiles={setFiles as any}
                labelIdle='Drag & Drop your files'
                server={SERVER}
            />
        </div>
    )
}
```