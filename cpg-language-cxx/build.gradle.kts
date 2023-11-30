/*
 * Copyright (c) 2023, Fraunhofer AISEC. All rights reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *                    $$$$$$\  $$$$$$$\   $$$$$$\
 *                   $$  __$$\ $$  __$$\ $$  __$$\
 *                   $$ /  \__|$$ |  $$ |$$ /  \__|
 *                   $$ |      $$$$$$$  |$$ |$$$$\
 *                   $$ |      $$  ____/ $$ |\_$$ |
 *                   $$ |  $$\ $$ |      $$ |  $$ |
 *                   \$$$$$   |$$ |      \$$$$$   |
 *                    \______/ \__|       \______/
 *
 */
plugins {
    id("cpg.frontend-conventions")
}

publishing {
    publications {
        named<MavenPublication>("cpg-language-cxx") {
            pom {
                artifactId = "cpg-language-cxx"
                name.set("Code Property Graph - C/C++ Frontend")
                description.set("A C/C++ language frontend for the CPG")
            }
        }
    }
}

dependencies {
    // Eclipse dependencies
    api(libs.eclipse.runtime) {
        // For some reason, this group name is wrong
        exclude("org.osgi.service", "org.osgi.service.prefs")
    }
    api(libs.osgi.service)
    api(libs.icu4j)

    // CDT
    api(libs.eclipse.cdt.core)

    testImplementation(libs.junit.params)
}
