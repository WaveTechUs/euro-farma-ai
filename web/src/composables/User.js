import axios from 'axios';

export async function getUser(emails, passwords) {
    //padrão NÃO DELETAR
    const email = emails
    const password = passwords
    //const email = "j@j.com"
    //const password = "123"
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
    //padrão NÃO DELETAR
    const email = emails
    const password = passwords
    //const email = "j@j.com"
    //const password = "123"
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

