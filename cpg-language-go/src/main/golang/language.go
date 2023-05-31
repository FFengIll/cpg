/*
 * Copyright (c) 2021, Fraunhofer AISEC. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
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
package cpg

import (
	"tekao.net/jnigi"
)

type Language Node

const TranslationContextClass = CPGPackage + "/TranslationContext"

const FrontendsPackage = CPGPackage + "/frontends"
const GolangPackage = FrontendsPackage + "/golang"
const LanguageClass = FrontendsPackage + "/Language"
const LanguageFrontendClass = FrontendsPackage + "/LanguageFrontend"
const GoLanguageFrontendClass = GolangPackage + "/GoLanguageFrontend"

func (l *Language) ConvertToGo(o *jnigi.ObjectRef) error {
	*l = (Language)(*o)
	return nil
}

func (l *Language) ConvertToJava() (obj *jnigi.ObjectRef, err error) {
	return (*jnigi.ObjectRef)(l), nil
}

func (l *Language) GetClassName() string {
	return LanguageClass
}

func (l *Language) IsArray() bool {
	return false
}

func (ctx *TranslationContext) ConvertToGo(o *jnigi.ObjectRef) error {
	*ctx = (TranslationContext)(*o)
	return nil
}

func (ctx *TranslationContext) ConvertToJava() (obj *jnigi.ObjectRef, err error) {
	return (*jnigi.ObjectRef)(ctx), nil
}

func (*TranslationContext) GetClassName() string {
	return TranslationContextClass
}

func (*TranslationContext) IsArray() bool {
	return false
}
