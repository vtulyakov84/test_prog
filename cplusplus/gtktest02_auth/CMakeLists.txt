cmake_minimum_required(VERSION 3.10)
project(MyApp)

set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)

# Find required packages
find_package(PkgConfig REQUIRED)
pkg_check_modules(GTKMM REQUIRED gtkmm-3.0)
pkg_check_modules(LIBXML++ REQUIRED libxml++-2.6)
pkg_check_modules(LIBCURL REQUIRED libcurl)

# Include directories
include_directories(${GTKMM_INCLUDE_DIRS} ${LIBXML++_INCLUDE_DIRS} ${LIBCURL_INCLUDE_DIRS})
link_directories(${GTKMM_LIBRARY_DIRS} ${LIBXML++_LIBRARY_DIRS} ${LIBCURL_LIBRARY_DIRS})

# Source files
set(SOURCES
    src/main.cpp
    src/auth_window.cpp
    src/main_window.cpp
    src/rest_client.cpp
)

# Executable
add_executable(${PROJECT_NAME} ${SOURCES})

# Link libraries
target_link_libraries(${PROJECT_NAME}
    ${GTKMM_LIBRARIES}
    ${LIBXML++_LIBRARIES}
    ${LIBCURL_LIBRARIES}
)

# Copy UI files to build directory
file(COPY ${CMAKE_CURRENT_SOURCE_DIR}/ui/ DESTINATION ${CMAKE_CURRENT_BINARY_DIR}/ui)

# Set working directory to build dir for UI files
set_target_properties(${PROJECT_NAME} PROPERTIES
    RUNTIME_OUTPUT_DIRECTORY ${CMAKE_CURRENT_BINARY_DIR}
)