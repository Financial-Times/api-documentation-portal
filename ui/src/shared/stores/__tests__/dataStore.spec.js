jest.unmock('../dataStore')
import dataStore from '../dataStore';

describe('Data Store', () => {
  it('Should return the initial loading state', () => {
    const isLoading = dataStore.isLoading();
    expect(isLoading).toEqual(false);
  });

  it('Should return the updated loading state', () => {
    dataStore.loading = true;
    const isLoading = dataStore.isLoading();
    expect(isLoading).toEqual(true);
  });

  it('Should return the initial data store', () => {
    const data = dataStore.getCycles();
    expect(data).toHaveLength(0);
  });

  it('Should return the updated data store', () => {
    const mockData = [{foo:'bar'}]
    dataStore.cycles = mockData
    const data = dataStore.getCycles();

    expect(data).toHaveLength(1);
    expect(data[0]).toEqual({'foo':'bar'});
  });
});
