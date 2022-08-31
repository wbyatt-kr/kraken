// in src/App.js
import * as React from "react";
import axios from 'axios';
import {
  Admin,
  Resource,
  Edit,
  List,
  TextField,
  ReferenceField,
  Datagrid,
  ReferenceInput,
  Create,
  SimpleForm,
  TextInput,
  required,
  DeleteButton,
  SelectInput
} from 'react-admin';

const api = axios.create({
  baseURL: 'http://localhost:8081',
})

const lowerKeys = (obj) => {
  const newObj = {};

  Object.keys(obj).forEach(key => {
    newObj[key.toLowerCase()] = obj[key];
  });

  return newObj;
}

const getMany = (resource) => api.get(resource)
  .then(({ data }) => ({ data: (data[resource] || []).map(lowerKeys), total: (data[resource] || []).length }));

const dataProvider = {
  getList: getMany,
  getMany,
  getManyReference: (resource, { target, id }) => api.get(`${target}/${id}/${resource}`)
    .then(({ data }) => ({ data: data.map(lowerKeys), total: data.length })),
  getOne: (resource, { id }) => api.get(`${resource}/${id}`)
    .then(({ data }) => ({ data: lowerKeys(data[resource]) })),
  create: (resource, params) => api.post(resource, params.data),
  delete: (resource, params) => api.delete(`${resource}/${params.id}`),
  update: (resource, params) => api.put(`${resource}/${params.id}`, params.data),
}



const RoutesList = () => (
  <List>
    <Datagrid rowClick="edit">
      <TextField source="path" />
      <ReferenceField label="Backend" source="serviceid" reference="services" />
      <DeleteButton />
    </Datagrid>
  </List>
)

const RouteForm = () =>
  <SimpleForm>
    <TextInput source="path" validate={[required()]} />
    <ReferenceInput source="serviceid" reference="services">
      <SelectInput optionText="backend" label="Backend" />
    </ReferenceInput>
  </SimpleForm>


const RouteCreate = () => (
  <Create>
    <RouteForm />
  </Create>
)

const RouteEdit = () =>
  <Edit>
    <RouteForm />
  </Edit>

const ServiceCreate = () => (
  <Create>
    <SimpleForm>
      <TextInput source="backend" validate={[required()]} fullWidth />
    </SimpleForm>
  </Create>
);

const ServicesList = () => (
  <List>
    <Datagrid>
      <TextField source="backend" />
      <DeleteButton />
    </Datagrid>
  </List>
)

const serviceRepresentation = (record) => record.backend;

const App = () =>
  <Admin dataProvider={dataProvider} >
    <Resource
      name="services"
      recordRepresentation={serviceRepresentation}
      list={ServicesList}
      create={ServiceCreate}
    />
    <Resource
      name="routes"
      list={RoutesList}
      create={RouteCreate}
      edit={RouteEdit}
    />
  </Admin >;

export default App;
