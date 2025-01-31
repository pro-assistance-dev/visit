// import { Module } from 'vuex';
//
// import UniqueUsersByHour from '@/classes/UniqueUsersByHour';
// import User from '@/classes/User';
// import UsersActivitiesTimesSum from '@/classes/UsersActivitiesTimesSum';
// import RootState from '@/services/interfaces/types';
// import getBaseDefaultState from '@/services/store/baseModule/baseIndex';
// import IBasicState from '@/services/store/baseModule/baseState';
//
// import actions from './actions';
// import getters from './getters';
// import mutations from './mutations';
//
// export interface State extends IBasicState<User> {
//   uniqueUsersByHour: UniqueUsersByHour[];
//   usersActivitiesTimesSum: UsersActivitiesTimesSum[];
// }
// export const getDefaultState = (): State => {
//   return {
//     ...getBaseDefaultState(User),
//     uniqueUsersByHour: [],
//     usersActivitiesTimesSum: [],
//   };
// };
//
// const state = getDefaultState();
// const namespaced = true;
//
// export const usersEventsActivities: Module<State, RootState> = {
//   namespaced,
//   state,
//   getters,
//   actions,
//   mutations,
// };
