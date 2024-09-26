import { getUser, postUser } from '@/composables/User';
import { useGoTo, useToast } from '@/composables/Utils';

export async function authUser(email, password) {

     const get = await getUser(email, password)
     if (get) {
         navigator.clipboard.writeText(get)
             .then(() => {
                 useToast().success('Seja bem vindo de volta!', {
                     timeout: 1500,
                     style: {
                         fontSize: '16px',
                         width: '250px',
                         height: '50px'
                     }
                 })
                 console.log(email, password)
                 console.log('Texto copiado com sucesso!');
                 useGoTo('/dashboard')

             })
             .catch((error) => {

                 console.error('Falha ao copiar o texto:', error);
             });
     }else{
         useToast().error('Deu erro no Login', {
                     timeout: 1500,
                     style: {
                         fontSize: '16px',
                         width: '250px',
                         height: '50px'
                     }
                 })
     }

 }
export async function authUserPost(email, password, nome) {
    

     const get = await postUser(email, password, nome)
     if (get) {
         navigator.clipboard.writeText(get)
             .then(() => {
                 useToast().success('Seja bem vindo de volta!', {
                     timeout: 1500,
                     style: {
                         fontSize: '16px',
                         width: '250px',
                         height: '50px'
                     }
                 })
                 console.log(email, password)
                 console.log('Texto copiado com sucesso!');
                 useGoTo('/dashboard')

             })
             .catch((error) => {

                 console.error('Falha ao copiar o texto:', error);
             });
     }else{
         useToast().error('Deu erro na criação', {
                     timeout: 1500,
                     style: {
                         fontSize: '16px',
                         width: '250px',
                         height: '50px'
                     }
                 })
     }

 }

