export const addProperty = async (materialId, propertyName) => {
    const respAddValue = await fetch('http://localhost:8080/addPropertyToMaterial', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            "material_id": materialId,
            "property_name": propertyName,
            "kind": "decimal"
        })
    })
}