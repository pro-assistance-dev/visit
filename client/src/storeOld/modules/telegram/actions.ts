// import { ActionTree } from 'vuex';
//
// import axiosInstance from '@/services/Axios';
// import RootState from '@/services/interfaces/types';
//
// import { State } from './index';
//
// const actions: ActionTree<State, RootState> = {
//   sendCallBackMessage: async (_, message: string): Promise<void> => {
//     if (!import.meta.env.VITE_APP_TG_TOKEN || import.meta.env.VITE_APP_TG_CHAT_ID) {
//       console.log('env tg');
//     }
//     await axiosInstance({
//       url: `https://api.telegram.org/bot${import.meta.env.VITE_APP_TG_TOKEN}/sendMessage?chat_id=${
//         import.meta.env.VITE_APP_TG_CHAT_ID
//       }&text=${message}`,
//       method: 'post',
//     });
//   },
// };
//
// export default actions;
