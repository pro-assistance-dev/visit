import Role from './Role';

export default class User {
  id?: string;
  email = '';
  password?: string = '';
  phone = '';
  lastName = '';
  firstName = '';
  patronName = '';
  company = '';
  post = '';
  privacy = false;
  role = new Role();
  roleId?: string;

  //
  fullName?: string;

  constructor(i?: User) {
    PF.H.Classes.BuildClass(this, i);
  }

  getName(): string {
    return this.getFullName();
  }

  getFullName(): string {
    return `${this.lastName} ${this.firstName} ${this.patronName}`;
  }
  isFullInfo(): boolean {
    return !!this.phone && !!this.lastName && !!this.firstName && !!this.patronName && !!this.company && !!this.post;
  }
}
