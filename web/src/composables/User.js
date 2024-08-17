import axios from 'axios';


export async function getUser() {
    try {
        const response = await axios.get('http://localhost:8080/user');
        console.log('Status:', response.status);
        console.log('Data:', response.data);
    } catch (error) {
        console.error('Error fetching user:', error.message);
    }
}

