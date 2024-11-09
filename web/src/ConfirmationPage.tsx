import {API_URL} from "./App.tsx";
import {useNavigate, useParams} from "react-router-dom";

export const ConfirmationPage = () => {
    const handleConfirm = async () => {
        const {token = ''} = useParams()
        const redirect = useNavigate()

        const response = await fetch(`${API_URL}/users/activate/${token}`, {
            method: "PUT"
        })

        if (response.ok) {
            redirect("/")
        } else {
            // handle error
            alert("Failed to confirm token")
        }

    }
    return (
        <div>
            <h1>Confirmation</h1>
            <button onClick={handleConfirm}>Click to confirm</button>
        </div>
    )
}