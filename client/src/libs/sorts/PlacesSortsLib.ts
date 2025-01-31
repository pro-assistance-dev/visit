import { log } from 'util';

import Place from '@/classes/Place';

const PlacesSortsLib = (() => {
  function byName(): SortModel {
    return SortModel.Create(Place, PF.H.Classes.GetPropertyName(Place).name);
  }
  return { byName };
})();

export default PlacesSortsLib;
