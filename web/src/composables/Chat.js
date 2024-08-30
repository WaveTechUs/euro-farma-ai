import axios from 'axios';
export async function PostAsk(ask, history) {
    try {
        const response = await axios.post('http://localhost:8080/gemini', 
        {
            Text:ask ,
            History:history
        });        
        console.log('Status:', response.status);
        console.log('Data:', response.data);
        return response.data

    }  catch (error) {
        console.error('Error during POST request:', error.message);
        if (error.response) {
            console.error('Status:', error.response.status);
            console.error('Data:', error.response.data);
        }
        throw error;
    }
}

