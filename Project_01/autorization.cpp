#include <iostream> 
#include <map> 
#include <vector> 
#include <string> 
#include <thread> 
#include <chrono> 
#include <limits> 
#include <httplib.h> 
#include <nlohmann/json.hpp> 

using json = nlohmann::json;

struct User {
    std::string Name;
    int64_t UserID;
    std::string AccessToken;
    std::vector<int> Access;
};

struct UserData {
    int64_t Id;
    std::string Name;
};

struct AuthenticateData {
    bool is_done;
    std::string code;
};

// Глобальная переменная для проверки, что пользователь дал доступ 
AuthenticateData authenticate{ false, "" };

// Данные GitHub приложения 
const std::string CLIENT_ID = "4997b657fb1f65f5d221";
const std::string CLIENT_SECRET = "1901a860dac9caaa98b9298dd2714ddc2a5116c7";

// Список пользователей изначально пуст 
std::map<int64_t, User> users;

// Функция для обработки запроса на /oauth 
void handleOauth(const httplib::Request& req, httplib::Response& res) {
    std::string responseHtml = "<html><body><h1>Вы НЕ аутентифицированы!</h1></body></html>";

    std::string code = req.get_param_value("code");
    if (!code.empty()) {
        authenticate.is_done = true;
        authenticate.code = code;
        responseHtml = "<html><body><h1>Вы аутентифицированы!</h1></body></html>";
    }

    res.set_content(responseHtml, "text/html");
}

// Функция для отправки POST-запроса 
httplib::Result sendPostRequest(const std::string& host, const std::string& path, const std::map<std::string, std::string>& params) {
    httplib::Client cli(host.c_str());
    httplib::Params httplibParams;

    for (const auto& param : params) {
        httplibParams.emplace(param.first, param.second);
    }

    auto res = cli.Post(path.c_str(), httplibParams);

    return res;
}

// Функция для получения токена доступа 
std::string getAccessToken(const std::string& code) {
    std::string host = "github.com";
    std::string path = "/login/oauth/access_token";

    std::map<std::string, std::string> params = {
        {"client_id", CLIENT_ID},
        {"client_secret", CLIENT_SECRET},
        {"code", code}
    };

    auto res = sendPostRequest(host, path, params);

    if (res && res->status == 200) {
        json responseJson = json::parse(res->body);
        return responseJson["access_token"].get<std::string>();
    }

    return "";
}

// Функция для получения информации о пользователе 
UserData getUserData(const std::string& accessToken) {
    httplib::Client cli("api.github.com");
    std::string path = "/user";

    auto res = cli.Get(path.c_str(), {
        {"Authorization", ("Bearer " + accessToken).c_str()}
        });

    if (res && res->status == 200) {
        UserData userData;
        json responseJson = json::parse(res->body);
        userData.Id = responseJson["id"].get<int64_t>();
        userData.Name = responseJson["name"].get<std::string>();
        return userData;
    }

    return UserData{ 0, "" };
}

// Запуск сервера 
void startServer() {
    httplib::Server svr;

    svr.Get("/oauth", handleOauth);
    svr.listen("localhost", 8080);
}

int main() {
    setlocale(LC_ALL, "RU");

    std::thread serverThread(startServer);

    std::string authURL = "https://github.com/login/oauth/authorize?client_id=" + CLIENT_ID;
    while (!authenticate.is_done) {
        std::cout << "Чтобы зайти, перейдите по ссылке:\n" << authURL << "\nи нажмите Enter\n";
        std::cin.ignore(std::numeric_limits<std::streamsize>::max(), '\n');
    }

    std::string accessToken = getAccessToken(authenticate.code);
    UserData userData = getUserData(accessToken);

    if (users.find(userData.Id) == users.end()) {
        users[userData.Id] = { userData.Name, userData.Id, accessToken, {13} };
    }

    const User& user = users[userData.Id];
    std::cout << "Добро пожаловать, " << user.Name << "\n";

    std::cout << "В какую зону хотите попасть? ";
    int area;
    std::cin >> area;

    if (std::find(user.Access.begin(), user.Access.end(), area) == user.Access.end()) {
        std::cout << "Нет доступа в эту зону\n";
    }
    else {
        std::cout << "Доступ получен\n";
    }
    serverThread.join();
    return 0;
}
