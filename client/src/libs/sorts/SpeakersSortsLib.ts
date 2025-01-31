import { log } from 'util';

import Speaker from '@/classes/Speaker';

const SpeakersSortsLib = (() => {
  function byFullName(order?: Orders): SortModel {
    return SortModel.Create(
      Speaker,
      PF.H.Classes.GetPropertyName(Speaker).fullName,
      order ? order : Orders.Asc,
      `По ФИО ${order === Orders.Asc ? '(вверх)' : '(вниз)'}`,
      order === Orders.Desc ? false : true
    );
  }
  return { byFullName };
})();

export default SpeakersSortsLib;
