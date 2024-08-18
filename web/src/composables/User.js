import axios from 'axios';
import { useGoTo } from '@/composables/Utils';

export async function getUser() {
    const email = "j@j.com"
    const password = "123"
    try {
        const response = await axios.post('http://localhost:8080/user/test/login', 
        {
            email: email,
            password: password
        });
        console.log('Status:', response.status);
        console.log('Data:', response.data);
    } catch (error) {
        console.error('Error fetching user:', error.message);
    }
    useGoTo('/dashboard')
}

