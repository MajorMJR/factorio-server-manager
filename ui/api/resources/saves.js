import client from "../client";

export default {
    list: async () => {
        const response = await client.get('/api/saves/list');
        return response.data;
    },
    delete: async (save) => {
        const response = await client.get(`/api/saves/rm/${save.name}`);
        return response.data;
    },
    defaultMapSettings: async () => {
        const response = await client.get('/api/saves/default-map-settings');
        return response.data;
    },
    defaultMapGenSettings: async () => {
        const response = await client.get('/api/saves/default-map-gen-settings');
        return response.data;
    },
    preview: async mapSettings => {
        const response = await client.post(`/api/saves/preview`, mapSettings);
        return response.data;
    },
    create: async (name, settings) => {
        const response = await client.post(`/api/saves/create/${name}`, settings);
        return response.data;
    },
    upload: async file => {
        let formData = new FormData();
        formData.append("savefile", file);

        const response = await client.post(`/api/saves/upload`, formData, {
            headers: {
                "Content-Type": "multipart/form-data"
            }
        });
        return response.data;
    },
    mods: async save => {
        const response = await client.post("/api/saves/mods", {
            saveFile: save
        });
        return response.data;
    }
}