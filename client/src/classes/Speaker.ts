import { AcademicDegree, AcademicRank, AcademicRelation } from '@/interfaces/AcademicEnums';

import Experience from './Experience';

export default class Speaker {
  id?: string;
  position = '';
  city = '';
  type = '';
  format = '';
  academicDegree?: AcademicDegree;
  academicRank?: AcademicRank;
  academicRelation?: AcademicRelation;
  human = new PF.C.Human();
  humanId?: string;

  @PF.H.Classes.GetClassConstructor(Experience)
  experiences: Experience[] = [];
  //
  fullName?: string;
  constructor(i?: Speaker) {
    PF.H.Classes.BuildClass(this, i);
  }

  static Create(): Speaker {
    const item = new Speaker();
    item.id = PF.H.Id.UUID();

    item.human.id = PF.H.Id.UUID();
    item.humanId = item.human.id;

    // item.name = "Новая локация"
    return item;
  }
  setAcademicDegree(item: AcademicDegree): void {
    this.academicDegree = item;
  }
  setAcademicRank(item: AcademicRank): void {
    this.academicRank = item;
  }
  setAcademicRelation(item: AcademicRelation): void {
    this.academicRelation = item;
  }

  static GetClassName(): string {
    return 'speaker';
  }

  getFullRegalias(): string {
    let r = this.position;
    if (this.academicDegree) {
      r += ', ' + this.academicDegree;
    }
    if (this.academicRank) {
      r += ', ' + this.academicRank;
    }
    return r;
  }

  addExperience(): Experience {
    const item = Experience.Create(this.id);
    this.experiences.push(item);
    return item;
  }

  removeExperience(id: string): void {
    PF.H.Classes.RemoveFromClassById(id, this.experiences, []);
  }
}
