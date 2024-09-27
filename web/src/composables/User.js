import axios from 'axios';

export async function getUser(emails, passwords) {
    const email = emails
    const password = passwords
    try {
        const response = await axios.post('http://localhost:8080/user/login', 
        {
            email: email,
            password: password
        },
        {
            withCredentials: true 
        }
        );
        console.log('Status:', response.status);
        console.log('Data:', response.data);
        return response.data
    } catch (error) {
        console.error('Error fetching user:', error.message);
    }
}
export async function postUser(emails, passwords, nome) {
    const email = emails
    const password = passwords
    try {
        const response = await axios.post('http://localhost:8080/user/register', 
        {
            name: nome,
            role: 'RH',
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

