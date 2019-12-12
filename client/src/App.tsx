import React, { useState } from 'react'
import { FilePond, File } from 'react-filepond'
import 'filepond/dist/filepond.min.css'

const SERVER = 'http://localhost:8010/upload'

export default function App() {
    const [files, setFiles] = useState<File[]>([])
    return (
        <div className='App'>
            <FilePond
                files={files}
                allowMultiple={true}
                onupdatefiles={setFiles as any}
                labelIdle='Drag & Drop your files or <span class="filepond--label-action">Browse</span>'
                server={SERVER}
            />
        </div>
    )
}
