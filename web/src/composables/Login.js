import { getUser } from '@/composables/User';
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
         useToast().error('Deu erro fera', {
                     timeout: 1500,
                     style: {
                         fontSize: '16px',
                         width: '250px',
                         height: '50px'
                     }
                 })
     }

 }
