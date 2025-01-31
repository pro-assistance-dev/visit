import Perfom from '@/classes/Perfom';

const PerfomsSortsLib = (() => {
  function byName(): SortModel {
    return SortModel.CreateSortModel(Perfom, PF.H.Classes.GetPropertyName(Perfom).name);
  }
  return { byName };
})();

export default PerfomsSortsLib;
