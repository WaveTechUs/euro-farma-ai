import axios from 'axios';
import { useGoTo } from '@/composables/Utils';

export async function getUser(emails, passwords) {
    //padrão NÃO DELETAR
    // const email = "j@j.com"
    // const password = "123"
    const email = emails
    const password = passwords
    try {
        const response = await axios.post('http://localhost:8080/user/test/login', 
        {
            email: email,
            password: password
        });
        console.log('Status:', response.status);
        console.log('Data:', response.data);
        return response.data
    } catch (error) {
        console.error('Error fetching user:', error.message);
    }
}

