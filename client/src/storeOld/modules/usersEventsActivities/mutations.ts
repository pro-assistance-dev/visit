// import { MutationTree } from 'vuex';
//
// import UniqueUsersByHour from '@/classes/UniqueUsersByHour';
// import User from '@/classes/User';
// import UsersActivitiesTimesSum from '@/classes/UsersActivitiesTimesSum';
// import getBaseMutations from '@/services/store/baseModule/baseMutations';
//
// import { State } from './index';
//
// const mutations: MutationTree<State> = {
//   ...getBaseMutations(User),
//   setUsersActivitiesTimesSum(state, items: UsersActivitiesTimesSum[]): void {
//     state.usersActivitiesTimesSum = items.map((i: UsersActivitiesTimesSum) => new UsersActivitiesTimesSum(i));
//   },
//   setUniqueUsersByHour(state, items: UniqueUsersByHour[]): void {
//     state.uniqueUsersByHour = items.map((i: UniqueUsersByHour) => new UniqueUsersByHour(i));
//   },
// };
//
// export default mutations;
