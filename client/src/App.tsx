import React, { useState } from 'react'
// Import React FilePond
import { FilePond, File } from 'react-filepond'

// Import FilePond styles
import 'filepond/dist/filepond.min.css'

const SERVER = 'http://localhost:8010/upload'

export default function App() {
    const [files, setFiles] = useState<File[]>([])
    const onProcessFile = (err, file) => {
        if (!err) {
            console.log(file)
        }
    }
    console.log(files.length && files[0].status)
    return (
        <div className='App'>
            <FilePond
                files={files}
                allowMultiple={true}
                onupdatefiles={setFiles as any}
                labelIdle='Drag & Drop your files or <span class="filepond--label-action">Browse</span>'
                server={SERVER}
                onprocessfile={onProcessFile}
            />
        </div>
    )
}
