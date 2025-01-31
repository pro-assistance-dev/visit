// import { ActionTree } from 'vuex';
//
// import M2M from '@/services/classes/M2M';
// import HttpClient from '@/services/HttpClient';
// import RootState from '@/services/interfaces/types';
//
// import { State } from './index';
// const httpClient = new HttpClient('m2m');
//
// const actions: ActionTree<State, RootState> = {
//   create: async (_, item: M2M): Promise<void> => {
//     await httpClient.post<M2M, undefined>({
//       query: 'create',
//       payload: item,
//       isFormData: true,
//     });
//   },
//   remove: async (_, item: M2M): Promise<void> => {
//     await httpClient.post<M2M, undefined>({
//       query: 'delete',
//       payload: item,
//       isFormData: true,
//     });
//   },
// };
//
// export default actions;
