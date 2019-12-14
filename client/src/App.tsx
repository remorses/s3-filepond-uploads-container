import React, { useState } from 'react'
import { FilePond, File } from 'react-filepond'
import 'filepond/dist/filepond.min.css'

const SERVER = 'http://localhost:8010/upload'

export default function App() {
    const [files, setFiles] = useState<File[]>([])
    const [urls, setUrls] = useState<string[]>([])
    return (
        <div className='App'>
            <FilePond
                files={files}
                allowMultiple={true}
                onupdatefiles={(files) => {
                    // setUrls(files.map((x) => x.serverId))
                    setFiles(files)
                }}
                onprocessfile={(err, file) => {
                    console.log(file.serverId)
                    setUrls([...urls, file.serverId])
                }}
                labelIdle='Drag & Drop your files or <span class="filepond--label-action">Browse</span>'
                server={SERVER}
            />
            <pre>{JSON.stringify(urls, null, 4)}</pre>
        </div>
    )
}
