import axios from 'axios';
export async function GetDataDash() {
    try {
        const response = await axios.get('http://localhost:8080/survey');
        return response.data
    } catch (error) {
        if (error.response) {
        }
        throw error;
    }
}
