import { RoleName } from '@/interfaces/RoleName';

export default class Role {
  id?: string;
  name: RoleName = RoleName.User;
  label = '';
  startPage = '';

  constructor(i?: Role) {
    PF.H.Classes.BuildClass(this, i);
  }
}
