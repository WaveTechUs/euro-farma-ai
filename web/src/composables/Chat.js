import axios from 'axios';
export async function PostAsk(ask, history) {
    try {
        const response = await axios.post('http://localhost:8080/gemini',
            {
                Text: ask,
                History: history
            });
        return response.data
    } catch (error) {
        if (error.response) {
            console.error('Status:', error.response.status);
            console.error('Data:', error.response.data);
        }
        throw error;
    }
}

