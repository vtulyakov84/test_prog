cmake_minimum_required(VERSION 3.10)
project(hello_gtk VERSION 1.0.0)

# Устанавливаем стандарт C++
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

# Находим GTK+ библиотеку
find_package(PkgConfig REQUIRED)
pkg_check_modules(GTK REQUIRED gtk+-3.0)

# Создаем исполняемый файл
add_executable(${PROJECT_NAME}
            src/main.cpp)

# Подключаем GTK+ к цели
target_link_libraries(${PROJECT_NAME}
            PRIVATE
                ${GTK_LIBRARIES})

target_include_directories(${PROJECT_NAME}
            PRIVATE
                ${GTK_INCLUDE_DIRS})

target_compile_options(${PROJECT_NAME}
            PRIVATE
                ${GTK_CFLAGS_OTHER})

# Для отображения путей при сборке (опционально)
message(STATUS "GTK+ include dirs: ${GTK_INCLUDE_DIRS}")
message(STATUS "GTK+ libraries: ${GTK_LIBRARIES}")