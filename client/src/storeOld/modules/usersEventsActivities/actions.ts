// import { ActionTree } from 'vuex';
//
// import Role from '@/classes/Role';
// import HttpClient from '@/services/HttpClient';
// import RootState from '@/services/interfaces/types';
// import getBaseActions from '@/services/store/baseModule/baseActions';
//
// import { State } from './index';
//
// const httpClient = new HttpClient('users-events-activities');
//
// const actions: ActionTree<State, RootState> = {
//   ...getBaseActions(httpClient),
//   ping: async ({ commit }): Promise<void> => {
//     await httpClient.get({ query: `ping` });
//   },
//   uniqueUsersByHour: async ({ commit }): Promise<void> => {
//     commit('setUniqueUsersByHour', await httpClient.get({ query: `unique-users-by-hour` }));
//   },
//   usersActivitiesTimesSums: async ({ commit }): Promise<void> => {
//     commit('setUsersActivitiesTimesSum', await httpClient.get({ query: `users-activities-times-sums` }));
//   },
// };
//
// export default actions;
