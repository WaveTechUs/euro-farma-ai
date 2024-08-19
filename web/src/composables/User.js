import axios from "axios";
const URL = 'http://localhost:8080/';

export const listUser =()=>{
    try {
        let resp = axios.get('http://localhost:8080/user')
        console.log(resp.data);       
        return resp.data

    } catch (error) {
        console.error(error)
    }
}

