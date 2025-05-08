db.filters.insertMany([
  {
    "name": "size",
    "allowedValues": ["39", "40", "41", "42"]
  },
  {
    "name": "color",
    "allowedValues": ["черный", "белый", "красный"]
  },
  {
    "name": "brand",
    "allowedValues": ["Nike", "Adidas", "Puma"]
  },
  {
    "name": "material",
    "allowedValues": ["кожа", "текстиль", "замша"]
  },
  {
    "name": "minPrice",
    "allowedValues": ["5000"]
  },
  {
    "name": "maxPrice",
    "allowedValues": ["15000"]
  }
])