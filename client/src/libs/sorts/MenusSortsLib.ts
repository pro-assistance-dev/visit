const MenusSortsLib = (() => {
  function byOrder(order?: Orders): SortModel {
    return SortModel.Create(Menu, PF.H.Classes.GetPropertyName(Menu).order, order ? order : Orders.Asc);
  }
  return {
    byOrder,
  };
})();

export default MenusSortsLib;
