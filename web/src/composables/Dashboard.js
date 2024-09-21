import axios from 'axios';
export async function GetDataDash() {
    try {
        const response = await axios.get('http://localhost:8080/survey');        
        console.log('Status:', response.status);
        console.log('Data:', response.data);
        return response.data

    }  catch (error) {
        console.error('Error during Get request:', error.message);
        if (error.response) {
            console.error('Status:', error.response.status);
            console.error('Data:', error.response.data);
        }
        throw error;
    }
}
