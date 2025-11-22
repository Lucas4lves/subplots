import { useState } from "react"

interface UpdateStatusBtnProps {
    status: boolean
    onStatusChange : (newStatus : boolean) => void
}

export const UpdateStatusBtn = ({status, onStatusChange} : UpdateStatusBtnProps) => {
    const [btnStatus, setBtnStatus] = useState(false)

    const update = (e : React.MouseEvent) => {
        setBtnStatus(!btnStatus)
    }

    return (
        <button onClick={update(e)}>
            
        </button>
    )
}